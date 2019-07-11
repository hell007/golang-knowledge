package services

import (
	"../dao"
	models "../models/system"
	db "../utils/datasource"
	"../utils/page"
)

type MenuService interface {
	GetAll() []models.Menu
	DynamicMenuTree(uid int64) []models.Menu
	GetMenusByRoleid(rid int64, page *page.Pagination) ([]models.Menu, int64, error)
}

type menuService struct {
	dao *dao.MenuDao
}

func NewMenuService() MenuService {
	return &menuService{
		dao: dao.NewMenuDao(db.MasterEngine()),
	}
}

func (s *menuService) GetAll() []models.Menu {
	return s.dao.GetAll()
}

func (s *menuService) DynamicMenuTree(uid int64) []models.Menu {
	return s.dao.DynamicMenuTree(uid)
}

func (s *menuService) GetMenusByRoleid(rid int64, page *page.Pagination) ([]models.Menu, int64, error) {
	return s.dao.GetMenusByRoleid(rid, page)
}
