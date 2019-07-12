/**
 * name: routes
 * author: jie
 * date: 2019-6-4
 * note: 路由控制
 */

package routes

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"../../framework/bootstrap"
	"../../framework/middleware"
	"../../framework/middleware/cors"
	"../../framework/services"
	"../controllers"
)

// Configure: registers the necessary routes to the app.

// mvc模式
func Configure(b *bootstrap.Bootstrapper) {

	fmt.Println("routes:==>定义路由")

	/* 定义路由 */
	main := b.Party("/", cors.Mycors()).AllowMethods(iris.MethodOptions)
	main.Use(middleware.ServeHTTP)

	// 首页模块
	home := main.Party("/")
	home.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	// 系统模块
	sys := main.Party("/sys")
	{
		//系统用户子模块
		user := mvc.New(sys.Party("/user"))
		userService := services.NewUserService()
		user.Register(userService)
		user.Handle(new(controllers.UserController))

		//系统角色子模块
		role := mvc.New(sys.Party("/role"))
		roleService := services.NewRoleService()
		role.Register(roleService)
		role.Handle(new(controllers.RoleController))

		//系统casbinrule子模块
		rule := mvc.New(sys.Party("/rule"))
		ruleService := services.NewRuleService()
		rule.Register(ruleService)
		rule.Handle(new(controllers.RuleController))
	}

	//user := mvc.New(main.Party("/user"))
	//userService := services.NewUserService()
	//user.Register(userService)
	//user.Handle(new(controllers.UserController))
}
