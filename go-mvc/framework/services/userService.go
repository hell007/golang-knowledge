/**
 * name: userService
 * author: jie
 * date: 2019-6-4
 * note:
 */

package services

import (
	"../dao"
	models "../models/system"
	"../utils/datasource"
	"../utils/page"
)

type UserService interface {
	//GetAll() []models.User
	List(name string, p *page.Pagination) ([]models.User, int64, error)
	Get(id int) *models.User
	GetUserByName(name string, user *models.User) (bool, error)
	GetRoleNameByRId(rid int) (string, error)
	Delete(ids []int) (int64, error)
	Update(user *models.User, columns []string) (int64, error)
	Create(user *models.User) (int64, error)
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.MasterEngine()),
	}
}

// func (s *userService) GetAll() []models.User {
// 	return s.dao.GetAll()
// }

func (s *userService) List(name string, p *page.Pagination) ([]models.User, int64, error) {
	return s.dao.List(name, p)
}

func (s *userService) Get(id int) *models.User {
	return s.dao.Get(id)
}

func (s *userService) GetUserByName(name string, user *models.User) (bool, error) {
	return s.dao.GetUserByName(name, user)
}

func (s *userService) GetRoleNameByRId(rid int) (string, error) {
	return s.dao.GetRoleNameByRId(rid)
}

func (s *userService) Update(user *models.User, columns []string) (int64, error) {
	return s.dao.Update(user, columns)
}

func (s *userService) Create(user *models.User) (int64, error) {
	return s.dao.Create(user)
}

func (s *userService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}
