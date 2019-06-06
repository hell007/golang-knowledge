/**
 * name: routes
 * author: jie
 * date: 2019-6-4
 * note: 路由控制
 */

package routes

import (
	"github.com/kataras/iris/mvc"

	"../../framework/bootstrap"
	"../../framework/services"
	"../controllers"
	//"../middleware"
)

// Configure: registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()

	user := mvc.New(b.Party("/user"))
	//user.Router.Use(middleware.BasicAuth)
	user.Register(userService)
	user.Handle(new(controllers.UserController))

	home := mvc.New(b.Party("/home"))
	home.Handle(new(controllers.HomeController))

	// b.Get("/", function(ctx iris.Context){
	///ctx.ViewData("Title", "Index Page")
	//ctx.View("index.html")
	// })

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
