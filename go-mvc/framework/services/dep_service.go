package services

import (
	"../dao"
	models "../models/system"
	db "../utils/datasource"
	"../utils/page"
)

type DepService interface {
	GetAll() []models.Dep
	List(p *page.Pagination) ([]models.Dep, int64, error)
	Get(id int) *models.Dep
	Delete(ids []int) (int64, error)
	Update(dep *models.Dep, columns []string) (int64, error)
	Create(dep *models.Dep) (int64, error)
}

type depService struct {
	dao *dao.DepDao
}

func NewDepService() DepService {
	return &depService{
		dao: dao.NewDepDao(db.MasterEngine()),
	}
}

func (s *depService) GetAll() []models.Dep {
	return s.dao.GetAll()
}

func (s *depService) List(p *page.Pagination) ([]models.Dep, int64, error) {
	return s.dao.List(p)
}

func (s *depService) Get(id int) *models.Dep {
	return s.dao.Get(id)
}

func (s *depService) Update(dep *models.Dep, columns []string) (int64, error) {
	return s.dao.Update(dep, columns)
}

func (s *depService) Create(dep *models.Dep) (int64, error) {
	return s.dao.Create(dep)
}

func (s *depService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}
