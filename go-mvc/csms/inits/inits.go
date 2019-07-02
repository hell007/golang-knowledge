package inits

import (
	"github.com/kataras/golog"

	system "../../framework/models/system"
	//"./parse"
	"./sys"
)

func init() {
	//parse.AppOtherParse()
	//parse.DBSettingParse()

	initRootUser()
}

func initRootUser() {
	// root is existed?
	if sys.CheckRootExit() {
		return
	}

	// create root user
	sys.CreateRoot()

	ok := sys.CreateSystemRole()
	if ok {
		addRoleMenu()
	}

}

func addRoleMenu() {
	// 添加role-menu关系
	rMenus := []*system.RoleMenu{
		{Rid: 68, Mid: 2},
		{Rid: 68, Mid: 3},
		{Rid: 68, Mid: 4},
		{Rid: 68, Mid: 5},
	}
	effect, err := system.CreateRoleMenu(rMenus...)
	if err != nil {
		golog.Fatalf("===> %d, %s", effect, err.Error())
	}
	golog.Infof("===> %s, %s", effect, err.Error())
}
