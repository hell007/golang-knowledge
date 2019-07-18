package dao

import (
	"github.com/go-xorm/xorm"

	models "../models/system"
	"../utils/page"
)

type RuleDao struct {
	engine *xorm.Engine
}

func NewRuleDao(engine *xorm.Engine) *RuleDao {
	return &RuleDao{
		engine: engine,
	}
}

// List
func (d *RuleDao) List(p *page.Pagination) ([]models.CasbinRule, int64, error) {

	list := make([]models.CasbinRule, 0)

	s := d.engine.Limit(p.Limit, p.Start)

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// Get
//func (d *RuleDao) Get(id int64) *models.CasbinRule {
//	data := &models.CasbinRule{Id: id}
//	ok, err := d.engine.Get(data)
//	if ok && err == nil {
//		return data
//	} else {
//		return nil
//	}
//}

// update
func (d *RuleDao) Update(rule *models.CasbinRule, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(rule.Id).MustCols(columns...).Update(rule)
	} else {
		effect, err = d.engine.Id(rule.Id).MustCols(columns...).Update(rule)
	}
	return effect, err
}

// insert
//func (d *RuleDao) Create(rule *models.CasbinRule) (int64, error) {
//	effect, err := d.engine.Insert(rule)
//	return effect, err
//}

// delete
func (d *RuleDao) Delete(ids []int64) (int64, error) {
	var (
		effect int64
		err    error
	)

	cr := new(models.CasbinRule)

	effect, err = d.engine.In("id", ids).Delete(cr)
	return effect, err
}
