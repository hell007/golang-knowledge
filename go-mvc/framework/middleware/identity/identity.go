package identity

import (
	"fmt"
	"github.com/kataras/iris"
	"time"

	"../../../framework/bootstrap"
)

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())
		ctx.Header("Server", "http://127.0.0.1:3000/")

		// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	fmt.Println("identity:==>定义header头")
	h := New(b)
	b.UseGlobal(h)
}
