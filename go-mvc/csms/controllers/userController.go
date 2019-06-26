/**
 * name: userController
 * author: jie
 * date: 2019-6-4
 * note:
 */

package controllers

import (
	//"fmt"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"

	"../../framework/middleware/jwt"
	"../../framework/models"
	"../../framework/services"
	//"../../framework/utils/encrypt"
	"../../framework/utils/page"
	"../../framework/utils/response"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

type UserToken struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"moblie"`
	RoleId int    `json:"roleId"`
	Token  string `json:"token"`
}

// user/login
func (c *UserController) PostLogin() {
	user := new(models.User)
	if err := c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", "", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.LoginFailur, nil)
		return
	}

	mUser := new(models.User)
	mUser.Name = user.Name
	has, err := c.Service.GetUserByName(mUser)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", user.Name, err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.LoginFailur, nil)
		return
	}

	// 用户名不正确
	if !has {
		response.Unauthorized(c.Ctx, response.UsernameFailur, nil)
		return
	}

	// 验证密码
	//ckPassword := encrypt.CheckPWD(user.Password, mUser.Password)
	//if !ckPassword {

	if user.Password != mUser.Password {
		response.Unauthorized(c.Ctx, response.PasswordFailur, nil)
		return
	}

	// 生成token
	token, err := jwt.GenerateToken(mUser)
	golog.Infof("用户[%s], 登录生成token [%s]", mUser.Name, token)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Name, err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenCreateFailur, nil)
		return
	}

	ut := UserToken{mUser.Id, mUser.Name, mUser.Email, mUser.Mobile, mUser.RoleId, token}
	response.Ok(c.Ctx, response.LoginSuccess, ut)
}

// user/list?pageNumber=1&pageSize=2&name=曹操
func (c *UserController) GetList() {
	name := c.Ctx.URLParam("name")
	p, _ := page.NewPagination(c.Ctx)
	list, total, err := c.Service.List(name, p)
	res := page.Result{}

	if err != nil {
		c.Ctx.Application().Logger().Errorf("UserController GetList失败。", err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	res.Total = total
	res.Rows = list
	response.Ok(c.Ctx, response.OptionSuccess, res)
}

// user/item?id=1
func (c *UserController) GetItem() {
	id, _ := c.Ctx.URLParamInt("id")
	user := c.Service.Get(id)

	if user != nil {
		response.Ok(c.Ctx, response.OptionSuccess, user)
		return
	} else {
		c.Ctx.Application().Logger().Errorf("UserController GetItem失败。%s", user)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}
}

// user/save
func (c *UserController) PostSave() {
	user := models.User{}

	if err := c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("UserController PostSave失败。", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	user.CreateTime = time.Now()
	rows, err2 := c.Service.Create(&user)

	if rows <= 0 || err2 != nil {
		c.Ctx.Application().Logger().Errorf("UserController PostSave失败。", err2.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}

// user/delete
func (c *UserController) PostDelete() {
	user := models.User{}

	if err := c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("UserController PostDelete失败。", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	rows, err := c.Service.Delete(user.Id)

	if rows <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("UserController GetDelete失败。", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}
