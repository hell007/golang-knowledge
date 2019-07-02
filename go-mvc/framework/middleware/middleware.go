package middleware

import (
	"strings"

	"github.com/kataras/iris/context"

	"../conf"
	"./casbin"
	"./jwt"
)

func ServeHTTP(ctx context.Context) {
	path := ctx.Path()

	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/assets") {
		ctx.Next()
		return
	}

	// jwt token拦截
	if !jwt.Serve(ctx) {
		return
	}

	// 系统菜单不进行权限拦截  mvc模式casbin出错了
	if !strings.Contains(path, "/sysMenu") {
		// casbin权限拦截
		ok := casbin.CheckPermissions(ctx)

		if !ok {
			return
		}
	}

	// Pass to real API
	ctx.Next()
}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
*/
func checkURL(reqPath string) bool {
	for _, v := range conf.AuthIgnores {
		if reqPath == v {
			return true
		}
	}

	return false
}
