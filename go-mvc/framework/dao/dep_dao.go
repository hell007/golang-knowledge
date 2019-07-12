package dao

import (
	"github.com/go-xorm/xorm"

	models "../models/system"
	"../utils/page"
)

type DepDao struct {
	engine *xorm.Engine
}

func NewDepDao(engine *xorm.Engine) *DepDao {
	return &DepDao{
		engine: engine,
	}
}

// GetAll
func (d *DepDao) GetAll() []models.Dep {
	datalist := make([]models.Dep, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// List
func (d *DepDao) List(p *page.Pagination) ([]models.Dep, int64, error) {

	list := make([]models.Dep, 0)

	s := d.engine.Limit(p.Limit, p.Start)

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// Get
func (d *DepDao) Get(id int) *models.Dep {
	data := &models.Dep{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

// update
func (d *DepDao) Update(dep *models.Dep, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(dep.Id).MustCols(columns...).Update(dep)
	} else {
		effect, err = d.engine.Id(dep.Id).MustCols(columns...).Update(dep)
	}
	return effect, err
}

// insert
func (d *DepDao) Create(dep *models.Dep) (int64, error) {
	effect, err := d.engine.Insert(dep)
	return effect, err
}

// delete
func (d *DepDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	dep := new(models.Dep)

	for _, v := range ids {
		i, err1 := d.engine.Id(v).Delete(dep)
		effect += i
		err = err1
	}
	return effect, err
}
