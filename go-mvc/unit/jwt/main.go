package main

import (
	//"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	//"github.com/kataras/iris"
	//"strings"
	"time"
)

// model
type User struct {
	Id          int      `json:"id"`
	Username    string   `json:"username" sql:"type:varchar(255), notnull, unique" binding:"required"`
	Password    string   `json:"-" sql:"type:varchar(255), notnull" binding:"required"`
	Fullname    string   `json:"fullname" sql:"type:varchar(255)"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "jie_secret" // 加盐
	ExpireTime = 1500         // token有效期
	StrToken   string
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	Id          int      `json:"d"`
	Username    string   `json:"username"`
	Fullname    string   `json:"fullname"`
	Permissions []string `json:"permissions"`
}

func (c *JWTClaims) SetExpiredAt(expiredAt int64) {
	c.ExpiresAt = expiredAt
}

func main() {
	GenerateToken()
	RefreshToken()
	//VerifyToken()
	ExpiresToken()
}

// 生成token
func GenerateToken() {
	user := User{
		Id:          1,
		Username:    "张三",
		Fullname:    "张大三",
		Permissions: []string{},
	}

	claims := JWTClaims{
		Id:          user.Id,
		Username:    user.Username,
		Fullname:    user.Fullname,
		Permissions: user.Permissions,
	}

	claims.IssuedAt = time.Now().Unix()
	claims.SetExpiredAt(time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix())

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(Secret))

	if err != nil {
		println(err)
		return
	}

	StrToken = token
	fmt.Println("GenerateToken--->" + StrToken + "\n")
}

// 刷新token
func RefreshToken() {
	token, err := jwt.ParseWithClaims(StrToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		golog.Errorf("RefreshToken解析token出错, %s", err)
		return
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		golog.Errorf("RefreshToken解析claims出错, %s", err)
		return
	}

	if err := token.Claims.Valid(); err != nil {
		golog.Errorf("RefreshToken验证token出错, %s", err)
		return
	}

	println("username=" + claims.Username)
	println(claims.IssuedAt)
	println(claims.ExpiresAt)

	return
	// user := User{
	// 	Id:          1,
	// 	Username:    "张三",
	// 	Fullname:    "张三",
	// 	Permissions: []string{},
	// }

	// claims.Fullname = user.Fullname
	// claims.Username = user.Username
	// claims.Permissions = user.Permissions
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := newToken.SignedString([]byte(Secret))
	if err != nil {
		println(err)
		return
	}
	fmt.Println("RefreshToken---->" + signedToken + "\n")
}

// 验证token是否有效
func VerifyToken() {
	token, err := jwt.ParseWithClaims(StrToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		println(err)
		return
	}

	if err := token.Claims.Valid(); err != nil {
		print(err)
		return
	}
	fmt.Println("VerifyToken ----> ok")
}

// 设置token失效
func ExpiresToken() {

	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(StrToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret), nil
		})

		if err != nil {
			fmt.Printf("err1===> %v", err)
		}

		if err := token.Claims.Valid(); err != nil {
			fmt.Printf("err2===> %v", err)
		}

		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Printf("err3===> %v", err)
		}
	})
}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

// 路由中间件
// func jwtAuthenticateMiddleware(ctx *iris.Context) {
// 	jwtObj := ctx.GetHeader("Authorization")
// 	if jwtObj == "" {
// 		ctx.AbortWithError(401, errors.New("Auth error, not find Autorization or Autorization is null"))
// 		return
// 	}

// 	jwtStr := strings.Split(jwtObj, "JWT ")[1]

// 	token, err := jwt.ParseWithClaims(jwtStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(configs.Default.Secret), nil
// 	})
// 	if err != nil {
// 		ctx.AbortWithError(401, err)
// 		return
// 	}
// 	claims, ok := token.Claims.(*JWTClaims)
// 	if !ok {
// 		ctx.AbortWithError(401, errors.New("test"))
// 		return
// 	}
// 	if err := token.Claims.Valid(); err != nil {
// 		ctx.AbortWithError(401, err)
// 		return
// 	}
// 	ctx.Set("USER_ID", claims.UserID)
// 	ctx.Next()
// }
