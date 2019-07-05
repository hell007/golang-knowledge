package response

import (
	"github.com/kataras/iris"
)

const (
	// key定义
	CODE  string = "code"
	STATE string = "state"
	MSG   string = "msg"
	DATA  string = "data"

	// msg define
	Success                  = "恭喜, 成功"
	OptionSuccess     string = "恭喜, 操作成功"
	OptionFailur      string = "抱歉, 操作失败"
	ParseParamsFailur string = "解析参数失败"

	RegisteSuccess  string = "恭喜, 注册用户成功"
	RegisteFailur   string = "注册失败"
	LoginSuccess    string = "恭喜, 登录成功"
	LoginFailur     string = "登录失败"
	LoginOutSuccess string = "恭喜, 退出成功"
	LoginOutFailur  string = "退出失败"

	UsernameFailur             string = "用户名错误"
	PasswordFailur             string = "密码错误"
	TokenCreateFailur          string = "生成token错误"
	TokenRefreshFailur         string = "刷新token错误"
	TokenExactFailur           string = "token不存在或header设置不正确"
	TokenExpire                string = "会话已过期"
	TokenParseFailur           string = "token解析错误"
	TokenParseFailurAndEmpty   string = "解析错误,token为空"
	TokenParseFailurAndInvalid string = "解析错误,token无效"
	NotFound                   string = "您请求的url不存在"
	PermissionsLess            string = "权限不足"

	RoleCreateFailur  string = "创建角色失败"
	RoleCreateSuccess string = "创建角色成功"

	// value define

)

// 200 define
func Ok_(ctx iris.Context, msg string) {
	Ok(ctx, msg, nil)
}

func Ok(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		CODE:  iris.StatusOK,
		STATE: true,
		MSG:   msg,
		DATA:  data,
	})
}

// 401 error define
func Unauthorized(ctx iris.Context, msg string, data interface{}) {
	unauthorized := iris.StatusUnauthorized

	ctx.StatusCode(unauthorized)
	ctx.JSON(iris.Map{
		CODE: unauthorized,
		MSG:  msg,
		DATA: data,
	})
}

// common error define
func Error(ctx iris.Context, status int, msg string, data interface{}) {
	ctx.StatusCode(status)
	ctx.JSON(iris.Map{
		CODE:  status,
		STATE: false,
		MSG:   msg,
		DATA:  data,
	})
}
