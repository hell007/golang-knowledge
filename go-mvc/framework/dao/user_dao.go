/**
 * name: userDao
 * author: jie
 * date: 2019-6-4
 * note:
 */

package dao

import (
	"fmt"
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

// GetUsersByRids
func (d *UserDao) GetUsersByRids(rids []int, page *page.Pagination) ([]models.User, int64, error) {
	users := make([]models.User, 0)
	s := d.engine.In("id", rids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&users)
	return users, count, err
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
func (d *UserDao) GetUserByName(name string, user *models.UserToken) (bool, error) {
	sql := fmt.Sprintf(`
SELECT USER.*, ROLE.ROLE_NAME AS ROLENAME, ROLE.ROLE_NOTE AS ROLENOTE FROM jie_user USER, jie_role ROLE 
WHERE USER.USERNAME = "%s" AND USER.ROLE_ID = ROLE.ID
`, name)
	//d.engine.Where("username = ?", name).Get(user)
	return d.engine.SQL(sql).Get(user)
}

// GetRoleNameById
func (d *UserDao) GetRoleNameByRId(rid int) (string, error) {
	var rolename string
	_, err := d.engine.Table("jie_role").Where("role_id = ?", rid).Cols("role_name").Get(&rolename)
	return rolename, err
}

// update
func (d *UserDao) Update(user *models.User, columns []string) (int64, error) {
	if columns != nil && len(columns) > 0 {
		effect, err := d.engine.Id(user.Id).MustCols(columns...).Update(user)
		return effect, err
	}

	effect, err := d.engine.Id(user.Id).Update(user)
	return effect, err
}

// insert
func (d *UserDao) Create(user *models.User) (int64, error) {
	effect, err := d.engine.Insert(user)
	return effect, err
}

// delete
func (d *UserDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	u := new(models.User)

	for _, v := range ids {
		i, err1 := d.engine.Id(v).Delete(u)
		effect += i
		err = err1
	}
	return effect, err
}
