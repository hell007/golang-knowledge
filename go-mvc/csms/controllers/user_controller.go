/**
 * name: userController
 * author: jie
 * date: 2019-6-4
 * note:
 */

package controllers

import (
	//"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"

	"../../framework/middleware/jwt"
	models "../../framework/models/system"
	"../../framework/services"
	"../../framework/utils/encrypt"
	"../../framework/utils/page"
	"../../framework/utils/response"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

// user/registe
func (c *UserController) PostRegiste() {
	var (
		err    error
		user   = new(models.User)
		effect int64
	)

	// 读取
	if err = c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		goto FAIL
	}

	// 新增
	user.Password = encrypt.AESEncrypt([]byte(user.Password))
	user.CreateTime = time.Now()

	effect, err = c.Service.Create(user)
	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.RegisteFailur, nil)
		return
	}

	response.Ok_(c.Ctx, response.RegisteSuccess)
	return

	// 失败处理
FAIL:
	response.Error(c.Ctx, iris.StatusBadRequest, response.RegisteFailur, nil)
	return
}

// user/login
func (c *UserController) PostLogin() {
	var (
		err        error
		user       = new(models.User)
		mUser      = new(models.User)
		ut         = new(models.UserToken) //需要返回的组装user
		ckPassword bool
		rolename   string
		token      string
	)

	if err = c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", "", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.LoginFailur, nil)
		return
	}

	mUser.Username = user.Username
	has, err := c.Service.GetUserByName(user.Username, mUser)

	if err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", user.Username, err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.LoginFailur, nil)
		return
	}

	// 用户名不正确
	if !has {
		response.Unauthorized(c.Ctx, response.UsernameFailur, nil)
		return
	}

	// 验证密码
	ckPassword = encrypt.CheckPWD(user.Password, mUser.Password)
	if !ckPassword {
		response.Unauthorized(c.Ctx, response.PasswordFailur, nil)
		return
	}

	// 查询角色
	rolename, err = c.Service.GetRoleNameByRId(mUser.RoleId)
	if err != nil {
		response.Unauthorized(c.Ctx, response.PermissionsLess, nil)
		return
	}

	// 组装前台需要数据
	ut.Id = mUser.Id
	ut.Username = mUser.Username
	ut.Email = mUser.Email
	ut.Mobile = mUser.Mobile
	ut.RoleId = mUser.RoleId
	ut.Rolename = rolename

	// 生成token
	token, err = jwt.GenerateToken(ut)
	golog.Infof("用户[%s], 登录生成token [%s]", mUser.Username, token)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Username, err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenCreateFailur, nil)
		return
	}

	ut.Token = token
	response.Ok(c.Ctx, response.LoginSuccess, ut)
}

// user/loginout token过期
func (c *UserController) GetLoginout() {

	response.Ok(c.Ctx, response.LoginOutSuccess, nil)
	return
}

// user/test 刷新token
func (c *UserController) GetTest() {
	var (
		err      error
		oldtoken string
		token    string
	)

	oldtoken, err = jwt.FromAuthHeader(c.Ctx)
	if err != nil {
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenParseFailur, nil)
		return
	}

	token, err = jwt.RefreshToken(oldtoken)
	if err != nil {
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenRefreshFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.LoginSuccess, token)
	return
}

// user/list?pageNumber=1&pageSize=2&name=曹操
func (c *UserController) GetList() {
	var (
		err      error
		username string
		p        *page.Pagination
		res      *page.Result
		list     []models.User
		total    int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	username = c.Ctx.URLParam("name")
	list, total, err = c.Service.List(username, p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("UserController GetList失败。", err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Total: total,
		Rows:  list,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
	return

	// 参数错误
FAIL:
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// user/item?id=1
func (c *UserController) GetItem() {
	var (
		err  error
		id   int
		user = new(models.User)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	user = c.Service.Get(id)
	if user != nil {
		response.Ok(c.Ctx, response.OptionSuccess, user)
		return
	} else {
		c.Ctx.Application().Logger().Errorf("UserController GetItem失败。%s", user)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	// 参数错误
FAIL:
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
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
	effect, err2 := c.Service.Create(&user)

	if effect <= 0 || err2 != nil {
		c.Ctx.Application().Logger().Errorf("UserController PostSave失败。", err2.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}

// user/delete?id=1,2
func (c *UserController) GetDelete() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int, 0)
		uid    int
		effect int64
	)

	id = c.Ctx.URLParam("id")
	idList = strings.Split(id, ",")
	if len(idList) == 0 {
		goto FAIL
	}

	for _, v := range idList {
		if v == "" {
			continue
		}

		uid, err = strconv.Atoi(v)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, uid)
	}

	effect, err = c.Service.Delete(ids)

	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("UserController PostDelete失败。", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}
