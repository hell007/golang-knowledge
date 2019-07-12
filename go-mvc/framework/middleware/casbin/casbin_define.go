package casbin

// RoleDefine 前端请求的结构体
type (
	RoleDefine struct {
		// 角色的标识等于casbin的sub，但角色需要加role_前缀
		Sub string `json:"sub"`
		// 对应casbin model的定义
		Obj      string `json:"obj"`
		Act      string `json:"act"`
		Suf      string `json:"suf"`
		RoleName string `json:"roleName"`
	}

	// GroupDefine 用户所属角色组
	GroupDefine struct {
		UID string   `json:"uid"` //uid设计为角色名称（string）， 而不是用户的id，避免了role id user id的区分
		Sub []string `json:"sub"`
	}
)
