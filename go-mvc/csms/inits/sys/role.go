package sys

import (
	"github.com/kataras/golog"

	"../../../framework/middleware/casbin"
)

var (
	// 定义系统初始的角色
	Components = [][]string{
		{"admin", "/admin*", "GET|POST|DELETE|PUT", ".*"},
		{"user", "/user*", "GET|POST|DELETE|PUT", ".*"},
	}
)

// 创建系统默认角色
func CreateSystemRole() bool {
	e := casbin.GetEnforcer()

	for _, v := range Components {
		p := e.GetFilteredPolicy(0, v[0])
		if len(p) == 0 {
			if ok := e.AddPolicy(v); !ok {
				golog.Fatalf("初始化角色[%s]权限失败。%s", v)
			}
		}
	}
	return true
}
