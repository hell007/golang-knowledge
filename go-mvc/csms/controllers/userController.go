/**
 * name: userController
 * author: jie
 * date: 2019-6-4
 * note:
 */

package controllers

import (
	//"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"time"

	"../../framework/models"
	"../../framework/services"
	"../../framework/utils/page"
	"../../framework/utils/result"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

// /user
func (c *UserController) Get() mvc.Result {
	datalist := c.Service.GetAll()

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "用户管理",
			"Base":     "http://127.0.0.1:3000/",
			"Datalist": datalist,
		},
		//Layout: "layout.html",
	}
}

// /user/list?pageNumber=1&pageSize=2&name=曹操
func (c *UserController) GetList() {
	name := c.Ctx.URLParam("name")
	p, _ := page.NewPagination(c.Ctx)
	list, total, err := c.Service.List(name, p)
	result := result.Result{}
	res := page.Result{}

	if err == nil {
		res.Total = total
		res.Rows = list
		result.State = true
		result.Message = "ok"
		result.Data = res
	} else {
		result.State = false
		result.Message = "fail"
		c.Ctx.Application().Logger().Errorf("userController GetList:", err.Error())
	}
	c.Ctx.JSON(result)
}

// /user/item?id=1
func (c *UserController) GetItem() {
	id, _ := c.Ctx.URLParamInt("id")
	result := result.Result{}
	user := c.Service.Get(id)

	if user != nil {
		result.Data = user
		result.Message = "ok"
		result.State = true
	} else {
		result.State = false
		result.Message = "fail"
		c.Ctx.Application().Logger().Errorf("userController GetItem:", user)
	}
	c.Ctx.JSON(result)
}

// /user/save
func (c *UserController) PostSave() {
	user := models.User{}
	result := result.Result{}

	if err := c.Ctx.ReadJSON(&user); err != nil {
		//fmt.Printf("1-->%v\n", err)
		c.Ctx.Application().Logger().Errorf("userController PostSave:", err.Error())
	}

	user.CreateTime = time.Now()
	rows, err2 := c.Service.Create(&user)

	if rows <= 0 || err2 != nil {
		result.State = false
		result.Message = "fail"
		c.Ctx.Application().Logger().Errorf("userController PostSave:", err2.Error())
	} else {
		result.Message = "ok"
		result.State = true
	}

	c.Ctx.JSON(result)
}

// /user/delete?id=1
func (c *UserController) GetDelete() {
	id, _ := c.Ctx.URLParamInt("id")
	result := result.Result{}
	rows, err := c.Service.Delete(id)

	if rows <= 0 || err != nil {
		result.State = false
		result.Message = "fail"
		c.Ctx.Application().Logger().Errorf("userController GetDelete:", err)
	} else {
		result.Message = "ok"
		result.State = true
	}

	c.Ctx.JSON(result)
}
