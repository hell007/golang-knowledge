package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./", ".html")
	htmlEngine.Reload(false)
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("你好,新的一天。--from iris go.")
	})

	app.Get("/test", func(ctx iris.Context) {
		ctx.ViewData("Title", "单元测试")
		ctx.ViewData("Content", "你好iris! -- from template")
		ctx.View("test.html")
	})

	app.Run(iris.Addr(":3000"), iris.WithCharset("UTF-8"))
}
