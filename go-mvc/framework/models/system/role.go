package system

type Role struct {
	RoleId      int    `json:"roleId" xorm:"not null pk autoincr comment('角色id') unique INT(10)"`
	Pid         int    `json:"pid" xorm:"not null default 0000000001 comment('角色id的父级id') INT(10)"`
	RoleName    string `json:"roleName" xorm:"not null comment('角色名称') unique VARCHAR(20)"`
	Description string `json:"description" xorm:"comment('角色职责描述') VARCHAR(100)"`
	Status      int    `json:"status" xorm:"not null default 1 comment('状态：启用=1/停用=2') TINYINT(3)"`
}
