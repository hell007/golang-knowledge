package system

import (
	"time"
)

type User struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('系统id') unique INT(10)"`
	RoleId   int    `json:"roleId" xorm:"not null comment('角色') INT(10)"`
	Username string `json:"username" xorm:"not null comment('系统用户') unique VARCHAR(20)"`
	Password string `json:"password" xorm:"not null comment('密码') CHAR(32)"`
	Name     string `json:"name" xorm:"not null comment('姓名') unique VARCHAR(20)"`
	Enable   int    `json:"enable" xorm:"not null default 1 comment('状态：启用=1/停用=0') TINYINT(1)"`
	Salt     string `json:"salt" xorm:"comment('盐值') VARCHAR(64)"`
	Email    string `json:"email" xorm:"comment('邮箱') unique VARCHAR(50)"`
	Mobile   string `json:"mobile" xorm:"not null comment('手机号码') unique VARCHAR(11)"`
	Ip       string `json:"ip" xorm:"comment('登录ip') VARCHAR(20)"`
	//IsSuper    int       `json:"isSuper" xorm:"not null default 1 comment('超级管理者：是/否 ') TINYINT(3)"`
	//Status     int       `json:"status" xorm:"not null default 1 comment('状态：启用=1/停用=2') TINYINT(3)"`
	CreateTime time.Time `json:"createTime" xorm:"not null comment('创建时间') DATETIME"`
	LoginTime  time.Time `json:"loginTime" xorm:"comment('登录时间') DATETIME"`
}
