package dao

import (
	"github.com/go-xorm/xorm"

	models "../models/system"
	"../utils/page"
)

type RoleDao struct {
	engine *xorm.Engine
}

func NewRoleDao(engine *xorm.Engine) *RoleDao {
	return &RoleDao{
		engine: engine,
	}
}

// List
func (d *RoleDao) List(p *page.Pagination) ([]models.CasbinRule, int64, error) {

	list := make([]models.CasbinRule, 0)

	s := d.engine.Limit(p.Limit, p.Start)

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// Get
func (d *RoleDao) Get(id int64) *models.CasbinRule {
	data := &models.CasbinRule{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

// update
func (d *RoleDao) Update(role *models.CasbinRule, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(role.Id).MustCols(columns...).Update(role)
	} else {
		effect, err = d.engine.Id(role.Id).MustCols(columns...).Update(role)
	}
	return effect, err
}

// insert
func (d *RoleDao) Create(role *models.CasbinRule) (int64, error) {
	effect, err := d.engine.Insert(role)
	return effect, err
}

// delete
func (d *RoleDao) Delete(ids []int64) (int64, error) {
	var (
		effect int64
		err    error
	)

	cr := new(models.CasbinRule)

	for _, v := range ids {
		i, err1 := d.engine.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return effect, err
}
