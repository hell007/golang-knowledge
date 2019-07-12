package system

import (
	db "../../../framework/utils/datasource"
)

// RoleMenu  角色-菜单关联表
type RoleMenu struct {
	Id  int64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	Rid int64 `json:"rid" xorm:"not null comment('角色id') INT(10)"`
	Mid int64 `json:"mid" xorm:"not null comment('菜单id') INT(10)"`
}

// CreateRoleMenu 建立角色菜单
func CreateRoleMenu(roleMenu ...*RoleMenu) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(roleMenu)
}
