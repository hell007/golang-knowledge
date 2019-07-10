package controllers

import (
	"github.com/kataras/iris"
	"net/http"
	"strconv"
	"strings"

	"../../framework/middleware/casbin"
	models "../../framework/models/system"
	"../../framework/services"
	"../../framework/utils/page"
	"../../framework/utils/response"
)

type RoleController struct {
	Ctx      iris.Context
	Service  services.RoleService
	UService services.UserService
}

// role/list?pageNumber=1&pageSize=2
func (c *RoleController) GetList() {
	var (
		err   error
		p     *page.Pagination
		res   *page.Result
		list  []models.CasbinRule
		total int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	list, total, err = c.Service.List(p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("RoleController GetList出错：", err.Error())
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

// 创建规则
func (c *RoleController) PostCreate() {
	var (
		err  error
		rule = new(models.CasbinRule)
	)

	// 读取
	if err = c.Ctx.ReadJSON(&rule); err != nil {
		c.Ctx.Application().Logger().Errorf("RoleController PostCreate出错：", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	e := casbin.GetEnforcer()
	ok := e.AddPolicy(rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5)
	if !ok {
		response.Error(c.Ctx, http.StatusInternalServerError, response.OptionFailur, nil)
	}
	response.Ok_(c.Ctx, response.OptionSuccess)
}

// 保存规则
func (c *RoleController) PostSave() {

	role := models.CasbinRule{}

	// 读取
	if err := c.Ctx.ReadJSON(&role); err != nil {
		c.Ctx.Application().Logger().Errorf("RoleController PostSave出错：", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err2 := c.Service.Update(&role, nil)

	if effect <= 0 || err2 != nil {
		c.Ctx.Application().Logger().Errorf("RoleController PostSave出错：", err2.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}
	response.Ok_(c.Ctx, response.OptionSuccess)
}

// role/delete?id=1,2
func (c *RoleController) GetDelete() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int64, 0)
		uid    int64
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

		uid, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, uid)
	}

	effect, err = c.Service.Delete(ids)

	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("RoleController PostDelete出错：", err.Error())
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

// RelationUserRole 给用户指定角色
func (c *RoleController) PostRelationUserRole() {
	groupDef := new(casbin.GroupDefine)

	if err := c.Ctx.ReadJSON(groupDef); err != nil {
		response.Error(c.Ctx, http.StatusInternalServerError, response.OptionFailur, err.Error())
		return
	}

	// TODO 校验前端的角色是否正确，和数据库的所有角色比较
	ok := true
	e := casbin.GetEnforcer()

	for _, v := range groupDef.Sub {
		// 给目标用户添加角色
		if !e.AddGroupingPolicy(strconv.FormatInt(groupDef.UID, 10), v) {
			ok = false
		}
	}

	if !ok {
		response.Error(c.Ctx, http.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	response.Ok_(c.Ctx, response.OptionSuccess)
}

// RoleUserTable  角色用户查询
func (c *RoleController) GetRoleuserlist() {
	rKey := c.Ctx.URLParam("rKey")
	p, err := page.NewPagination(c.Ctx)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("GetRoleuserlist出错: %s", err.Error())
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	e := casbin.GetEnforcer()
	users, err := e.GetUsersForRole(rKey)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("获取角色关联的用户表错误: %s", err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	uids := make([]int, 0)
	for _, v := range users {
		id, err := strconv.Atoi(v)
		if err != nil {
			c.Ctx.Application().Logger().Errorf("GetRoleuserlist出错：%s", err.Error())
			response.Error(c.Ctx, iris.StatusInternalServerError, response.ParseParamsFailur, nil)
			return
		}
		uids = append(uids, id)
	}

	list, total, err := c.UService.GetUsersByIds(uids, p)

	if err != nil {
		c.Ctx.Application().Logger().Errorf("获取角色关联的用户表错误, %s", err.Error())
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res := &page.Result{
		Total: total,
		Rows:  list,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
	return
}

// RoleMenuTable 角色菜单查询
//func (c *RoleController) RoleMenuTable(rid int64) {
//
//	p, err := page.NewPagination(c.Ctx)
//	if err != nil {
//		c.Ctx.Application().Logger().Errorf("获取角色关联的菜单表错误: %s", err.Error())
//		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
//		return
//	}
//
//	menus, total, err := c.Service.GetMenusByRoleid(rid, p)
//	if err != nil {
//		c.Ctx.Application().Logger().Errorf("获取角色关联的菜单表错误, %s, %v", err.Error(), menus)
//		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, err.Error())
//		return
//	}
//
//	// 组装数据
//	res := &page.Result{
//		Total: total,
//		Rows:  menus,
//	}
//	response.Ok(c.Ctx, response.OptionSuccess, res)
//	return
//}
