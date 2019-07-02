/**
 * name: userDao
 * author: jie
 * date: 2019-6-4
 * note:
 */

package dao

import (
	//"fmt"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	models "../models/system"
	"../utils/page"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {

	// tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "jie_")
	// engine.SetTableMapper(tbMapper)

	return &UserDao{
		engine: engine,
	}
}

// GetAll
func (d *UserDao) GetAll() []models.User {
	datalist := make([]models.User, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// List
func (d *UserDao) List(name string, p *page.Pagination) ([]models.User, int64, error) {

	list := make([]models.User, 0)

	s := d.engine.Where("username like ?", "%"+name+"%").Limit(p.Limit, p.Start)
	if p.SortName != "" {
		switch p.SortOrder {
		case "asc":
			s.Asc(p.SortName)
		case "desc":
			s.Desc(p.SortName)
		}
	}

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// Get
func (d *UserDao) Get(id int) *models.User {
	data := &models.User{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

// GetUserByName
func (d *UserDao) GetUserByName(user *models.User) (bool, error) {
	return d.engine.Get(user)
}

// update
func (d *UserDao) Update(data *models.User, columns []string) (int64, error) {
	rows, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return rows, err
}

// insert
func (d *UserDao) Create(data *models.User) (int64, error) {
	rows, err := d.engine.Insert(data)
	return rows, err
}

// delete
func (d *UserDao) Delete(id int) (int64, error) {
	data := &models.User{Id: id}
	rows, err := d.engine.Id(data.Id).Delete(data)
	return rows, err
}
