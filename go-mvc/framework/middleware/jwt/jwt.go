package jwt

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	"../../conf"
	models "../../models/system"
	"../../utils/response"
)

type (
	errorHandler   func(context.Context, string)
	TokenExtractor func(context.Context) (string, error)
	Jwts           struct {
		Config Config
	}
)

var (
	jwts *Jwts
	lock sync.Mutex
)

// Serve the middleware's action
func Serve(ctx context.Context) bool {
	// 设置jwt
	ConfigJWT()

	if err := jwts.CheckJWT(ctx); err != nil {
		//response.Unauthorized(ctx, response.Token_failur, nil)
		//ctx.StopExecution()
		golog.Errorf("Check jwt error, %s", err)
		return false
	}

	return true
	// If everything ok then call next.
	//ctx.Next()
}

// 返回此客户端/请求的用户（&token）信息
func (m *Jwts) Get(ctx context.Context) *jwt.Token {
	return ctx.Values().Get(m.Config.ContextKey).(*jwt.Token)
}

// 自定义log Printf
func (m *Jwts) logf(format string, args ...interface{}) {
	if m.Config.Debug {
		log.Printf(format, args...)
	}
}

// jwts中间件配置
func ConfigJWT() {
	if jwts != nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if jwts != nil {
		return
	}

	c := Config{
		ContextKey: DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(conf.JWTSecret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, errMsg string) {
			response.Error(ctx, iris.StatusUnauthorized, errMsg, nil)
		},
		// 指定func用于提取请求中的token
		Extractor: FromAuthHeader,
		// if the token was expired, expiration error will be returned
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	// 变量得到jwts
	jwts = &Jwts{Config: c}
	//return &Jwts{Config: c}
}

// 检查token的主要方法
func (m *Jwts) CheckJWT(ctx context.Context) error {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	// Use the specified token extractor to extract a token from the request
	token, err := m.Config.Extractor(ctx)
	// If an error occurs, call the error handler and return an error
	if err != nil {
		m.logf("Error extracting JWT: %v", err)
		m.Config.ErrorHandler(ctx, response.TokenExactFailur)
		return fmt.Errorf("Error extracting token: %v", err)
	}

	// If the token is empty...
	if token == "" {
		// Check if it was required
		if m.Config.CredentialsOptional {
			m.logf("  No credentials found (CredentialsOptional=true)")
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil
		}

		m.logf("  Error: No credentials found (CredentialsOptional=false)")
		// If we get here, the required token is missing
		m.Config.ErrorHandler(ctx, response.TokenParseFailurAndEmpty)
		return fmt.Errorf(response.TokenParseFailurAndEmpty)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		m.logf("Error parsing token1: %v", err)
		m.Config.ErrorHandler(ctx, response.TokenExpire)
		return fmt.Errorf("Error parsing token2: %v", err)
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {

		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])

		m.logf("Error validating token algorithm: %s", message)
		m.Config.ErrorHandler(ctx, response.TokenParseFailur) // 算法错误
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		m.logf(response.TokenParseFailurAndInvalid)
		m.Config.ErrorHandler(ctx, response.TokenParseFailurAndInvalid)
		return fmt.Errorf(response.TokenParseFailurAndInvalid)
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				return fmt.Errorf(response.TokenExpire)
			}
		}
	}

	//m.logf("JWT: %v", parsedToken)

	// If we get here, everything worked and we can set the
	// user property in context.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}

// 断言
type JWTClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Rolename string `json:"rolename"`
	jwt.StandardClaims
}

func (c *JWTClaims) SetExpiredAt(expiredAt int64) {
	c.ExpiresAt = expiredAt
}

// 在登录成功的时候生成token
func GenerateToken(ut *models.UserToken) (string, error) {

	claims := JWTClaims{
		ut.Id,
		ut.Username,
		ut.Rolename,
		jwt.StandardClaims{
			Issuer: "iris-casbins-jwt",
		},
	}

	claims.IssuedAt = time.Now().Unix()
	claims.SetExpiredAt(time.Now().Add(time.Second * time.Duration(conf.JWTTimeout)).Unix())

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(conf.JWTSecret))

	return token, err
}

// 根据原先的token刷新token的过期时间
func RefreshToken(signedToken string) (string, error) {

	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		golog.Errorf("RefreshToken解析token出错, %s", err)
		return signedToken, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		golog.Errorf("RefreshToken解析claims出错, %s", err)
		return signedToken, err
	}

	if err := token.Claims.Valid(); err != nil {
		golog.Errorf("RefreshToken验证token出错, %s", err)
		return signedToken, err
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := newToken.SignedString([]byte(conf.JWTSecret))

	return refreshToken, err
}

// 解析token的信息为当前用户
func ParseToken(ctx context.Context) (*models.UserToken, bool) {
	mapClaims := (jwts.Get(ctx).Claims).(jwt.MapClaims)

	id, ok1 := mapClaims["id"].(float64)
	rolename, ok2 := mapClaims["rolename"].(string)

	if !ok1 || !ok2 {
		response.Error(ctx, iris.StatusInternalServerError, response.TokenParseFailur, nil)
		return nil, false
	}

	ut := models.UserToken{
		Id:       int(id),
		Rolename: rolename,
	}

	return &ut, true
}

// 以下的方法都是从url获取token

// 来自授权头的JWT令牌
func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return "", nil // No error, just no token
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// 从指定的查询字符串参数中提取标记
func FromParameter(param string) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

// 运行多个令牌提取程序并获取它找到的第一个令牌
func FromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}
