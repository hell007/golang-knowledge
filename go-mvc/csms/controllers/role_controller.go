package controllers

import (
	"github.com/kataras/iris"
	"strconv"
	"strings"

	models "../../framework/models/system"
	"../../framework/services"
	"../../framework/utils/page"
	"../../framework/utils/response"
)

type RoleController struct {
	Ctx     iris.Context
	Service services.RoleService
}

// role/list?pageNumber=1&pageSize=2
func (c *RoleController) GetList() {
	var (
		err   error
		p     *page.Pagination
		res   *page.Result
		list  []models.Role
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

// role/delete?id=1,2
func (c *RoleController) GetDelete() {
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
