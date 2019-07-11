package dao

import (
	models "../models/system"
	"../utils/page"
	"fmt"
	"github.com/go-xorm/xorm"
)

type MenuDao struct {
	engine *xorm.Engine
}

func NewMenuDao(engine *xorm.Engine) *MenuDao {
	return &MenuDao{
		engine: engine,
	}
}

func (d *MenuDao) GetAll() []models.Menu {
	datalist := make([]models.Menu, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// DynamicMenuTree
func (d *MenuDao) DynamicMenuTree(uid int64) []models.Menu {

	sql := fmt.Sprintf(`
SELECT
	m1.id, m1.path, m1.modular, m1.component, m1.icon, m1.name, m1.require_auth,
	m2.id AS id2,
	m2.modular AS modular2,
	m2.component AS component2,
	m2.icon AS icon2,
	m2.keep_alive AS keep_alive2,
	m2.name AS name2,
	m2.path AS path2,
	m2.require_auth AS require_auth2
FROM jie_menu m1, jie_menu m2
WHERE m1.id = m2.parent_id
	AND m1.id != 1
	AND m2.id IN 
(
		SELECT rm.mid
		FROM jie_role_menu rm WHERE rm.rid in
		(
			SELECT id FROM jie_casbin_rule 
			WHERE 
			v2 <> 'ANY' AND 
			v0 in 
			(
				SELECT v1 FROM jie_casbin_rule WHERE v0=%d
			)
		)
)
AND m2.enabled=true order by m1.id, m2.id
`, uid)

	menuTree := make([]models.MenuTreeGroup, 0)
	d.engine.SQL(sql).Find(&menuTree)

	menus := make([]models.Menu, 0) // 菜单树

	if len(menuTree) > 0 {
		parentMenu := menuTree[0].Menu           // 父级菜单
		childMenus := make([]models.Children, 0) // 所有的子菜单
		for _, v := range menuTree {
			childMenus = append(childMenus, v.Children)
		}
		parentMenu.Children = childMenus

		menus = append(menus, parentMenu)
	}
	return menus
}

func (d *MenuDao) GetMenusByRoleid(rid int64, page *page.Pagination) ([]models.Menu, int64, error) {
	sql := fmt.Sprintf(`
SELECT * FROM jie_menu
WHERE id in
(
SELECT mid FROM jie_role_menu WHERE rid=%d
)
`, rid)

	sql2 := fmt.Sprintf(`
SELECT COUNT(*) FROM jie_role_menu WHERE rid=%d
`, rid)

	count, _ := d.engine.SQL(sql2).Count()

	menus := make([]models.Menu, 0)

	s := d.engine.Limit(page.Limit, page.Start)

	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			s.Asc(page.SortName)
		case "desc":
			s.Desc(page.SortName)
		}
	}

	err := s.SQL(sql).Find(&menus)

	return menus, count, err
}
