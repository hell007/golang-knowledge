/**
 * name: csms main
 * author: jie
 * date: 2019-6-4
 * note: csms入口
 */

package main

import (
	"../framework/bootstrap"
	//"./middleware/casbin"
	"./middleware/identity"
	"./routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("csms-system", "jie")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":3000")
}
