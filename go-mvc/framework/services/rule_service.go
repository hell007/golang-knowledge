package services

import (
	"../dao"
	models "../models/system"
	db "../utils/datasource"
	"../utils/page"
)

type RuleService interface {
	List(p *page.Pagination) ([]models.CasbinRule, int64, error)
	Update(rule *models.CasbinRule, columns []string) (int64, error)
	Delete(ids []int64) (int64, error)
}

type ruleService struct {
	dao *dao.RuleDao
}

func NewRuleService() RuleService {
	return &ruleService{
		dao: dao.NewRuleDao(db.MasterEngine()),
	}
}

func (s *ruleService) List(p *page.Pagination) ([]models.CasbinRule, int64, error) {
	return s.dao.List(p)
}

func (s *ruleService) Update(rule *models.CasbinRule, columns []string) (int64, error) {
	return s.dao.Update(rule, columns)
}

func (s *ruleService) Delete(ids []int64) (int64, error) {
	return s.dao.Delete(ids)
}
