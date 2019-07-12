package sys

import (
	"strconv"
	"time"

	"github.com/kataras/golog"

	"../../../framework/middleware/casbin"
	"../../../framework/models/system"
	db "../../../framework/utils/datasource"
	"../../../framework/utils/encrypt"
)

const (
	username = "root"
	password = "123456"
)

// 检查超级用户是否存在
func CheckRootExit() bool {
	e := db.MasterEngine()
	// root is existed?
	exit, err := e.Exist(&system.User{Username: username})
	if err != nil {
		golog.Fatalf("===>When check Root User is exited? happened error. %s", err.Error())
	}
	if exit {
		golog.Info("===>Root User is existed.")

		// 初始化rbac_model
		r := system.User{Username: username}

		if exit, _ := e.Get(&r); exit {
			casbin.SetRbacModel(strconv.Itoa(r.Id))
			CreateSystemRole()
		}
	}
	return exit
}

// CreateRoot 建立root用户
func CreateRoot() {
	newRoot := system.User{
		Username:   username,
		Password:   encrypt.AESEncrypt([]byte(password)),
		RoleId:     1,
		Enable:     1,
		Email:      "root@sina.com",
		Mobile:     "13888888888",
		CreateTime: time.Now(),
	}

	e := db.MasterEngine()
	if _, err := e.Insert(&newRoot); err != nil {
		golog.Fatalf("===>When create Root User happened error. %s", err.Error())
	}
	rooId := strconv.Itoa(newRoot.Id)
	casbin.SetRbacModel(rooId)

	addAllpolicy(rooId)
}

func addAllpolicy(rootId string) {
	rootId = "superadmin"
	// add policy for root
	e := casbin.GetEnforcer()
	p := e.AddPolicy(rootId, "/*", "ANY", ".*", "", "", "", "", "", "超级管理员")
	if !p {
		golog.Fatalf("初始化用户[%s]权限失败.", username)
	}

	for _, v := range Components {
		e.AddGroupingPolicy(rootId, v[0])
	}
}
