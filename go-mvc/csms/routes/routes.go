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

// api
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

	// 用户API模块
	user := mvc.New(main.Party("/user"))
	userService := services.NewUserService()
	user.Register(userService)
	user.Handle(new(controllers.UserController))

}

// mvc 模式
// func Configure(b *bootstrap.Bootstrapper) {
// 	userService := services.NewUserService()

// 	user := mvc.New(b.Party("/user"))
// 	user.Router.Use(middleware.BasicAuth)
// 	user.Register(userService)
// 	user.Handle(new(controllers.UserController))

// 	home := mvc.New(b.Party("/home"))
// 	home.Handle(new(controllers.HomeController))

// 	//b.Get("/follower/{id:long}", GetFollowerHandler)
// }
