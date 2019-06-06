/**
 * name: userController
 * author: jie
 * date: 2019-6-4
 * note:
 */

package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	//"../../framework/models"
	//"../../framework/services"
)

type HomeController struct {
	Ctx iris.Context
	//Service services.UserService
}

func (c *HomeController) Get() mvc.Result {
	//datalist := c.Service.GetAll()

	// set the model and render the view template.
	return mvc.View{
		Name: "home/index.html",
		Data: iris.Map{
			"Title": "首页",
		},
		Layout: "layout.html",
	}
}
