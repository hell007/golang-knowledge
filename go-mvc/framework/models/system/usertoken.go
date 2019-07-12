package system

type UserToken struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"moblie"`
	Rolename string `json:"rolename"`
	Rolenote string `json:"rolenote"`
	Token    string `json:"token"`
}
