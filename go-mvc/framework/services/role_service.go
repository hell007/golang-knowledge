package services

import (
	"../dao"
	models "../models/system"
	db "../utils/datasource"
	"../utils/page"
)

type RoleService interface {
	List(p *page.Pagination) ([]models.CasbinRule, int64, error)
	Update(role *models.CasbinRule, columns []string) (int64, error)
	Delete(ids []int64) (int64, error)
}

type roleService struct {
	dao *dao.RoleDao
}

func NewRoleService() RoleService {
	return &roleService{
		dao: dao.NewRoleDao(db.MasterEngine()),
	}
}

func (s *roleService) List(p *page.Pagination) ([]models.CasbinRule, int64, error) {
	return s.dao.List(p)
}

func (s *roleService) Update(role *models.CasbinRule, columns []string) (int64, error) {
	return s.dao.Update(role, columns)
}

func (s *roleService) Delete(ids []int64) (int64, error) {
	return s.dao.Delete(ids)
}
