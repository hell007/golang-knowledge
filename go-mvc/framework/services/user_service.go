/**
 * name: userService
 * author: jie
 * date: 2019-6-22
 * note:
 */

package services

import (
	"../dao"
	models "../models/system"
	db "../utils/datasource"
	"../utils/page"
)

type UserService interface {
	GetAll() []models.User
	List(name string, status int, p *page.Pagination) ([]models.User, int64, error)
	GetUsersByRids(rids []int, page *page.Pagination) ([]models.User, int64, error)
	Get(id int) *models.User
	GetUserByName(name string, user *models.UserToken) (bool, error)
	Update(user *models.User, columns []string) (int64, error)
	Create(user *models.User) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(db.MasterEngine()),
	}
}

func (s *userService) GetAll() []models.User {
	return s.dao.GetAll()
}

func (s *userService) List(name string, status int, p *page.Pagination) ([]models.User, int64, error) {
	return s.dao.List(name, status, p)
}

func (s *userService) Get(id int) *models.User {
	return s.dao.Get(id)
}

func (s *userService) GetUserByName(name string, user *models.UserToken) (bool, error) {
	return s.dao.GetUserByName(name, user)
}

func (s *userService) GetUsersByRids(rids []int, p *page.Pagination) ([]models.User, int64, error) {
	return s.dao.GetUsersByRids(rids, p)
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

func (s *userService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}
