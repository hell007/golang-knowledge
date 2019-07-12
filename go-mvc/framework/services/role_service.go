package services

import (
	"../dao"
	models "../models/system"
	db "../utils/datasource"
	"../utils/page"
)

type RoleService interface {
	GetAll() []models.Role
	List(p *page.Pagination) ([]models.Role, int64, error)
	Get(id int) *models.Role
	Update(role *models.Role, columns []string) (int64, error)
	Create(role *models.Role) (int64, error)
	Delete(ids []int) (int64, error)
}

type roleService struct {
	dao *dao.RoleDao
}

func NewRoleService() RoleService {
	return &roleService{
		dao: dao.NewRoleDao(db.MasterEngine()),
	}
}

func (s *roleService) GetAll() []models.Role {
	return s.dao.GetAll()
}

func (s *roleService) List(p *page.Pagination) ([]models.Role, int64, error) {
	return s.dao.List(p)
}

func (s *roleService) Get(id int) *models.Role {
	return s.dao.Get(id)
}

func (s *roleService) Update(role *models.Role, columns []string) (int64, error) {
	return s.dao.Update(role, columns)
}

func (s *roleService) Create(role *models.Role) (int64, error) {
	return s.dao.Create(role)
}

func (s *roleService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}
