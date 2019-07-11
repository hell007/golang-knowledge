package system

type UserToken struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"moblie"`
	RoleId   int    `json:"roleId"`
	Rolename string `json:"rolename"`
	Token    string `json:"token"`
}
