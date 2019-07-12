package system

type DepUser struct {
	Id    int `json:"id" xorm:"not null pk autoincr comment('id') INT(10)"`
	DepIp int `json:"depip" xorm:"comment('部门id') INT(10)"`
	Uid   int `json:"uid" xorm:"comment('用户id') INT(10)"`
}
