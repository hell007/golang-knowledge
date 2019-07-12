package system

type Dep struct {
	Id      int    `json:"id" xorm:"not null pk autoincr comment('部门id') unique INT(10)"`
	DepName string `json:"depname" xorm:"not null comment('部门名称') VARCHAR(64)"`
	Leader  string `json:"leader" xorm:"not null comment('部门领导人uid') VARCHAR(64)"`
	Tell    string `json:"tell" xorm:"comment('部门电话') VARCHAR(20)"`
	Email   string `json:"email" xorm:"comment('部门邮箱') VARCHAR(64)"`
	Status  int    `json:"status" xorm:"not null comment('0:可用 否则:不可用') TINYINT(1)"` //0:可用 否则:不可用
	Pid     int    `json:"pid" xorm:"not null INT(10)"`
	DepNote string `json:"depnote" xorm:"default '' comment('部门描述') VARCHAR(255)"`
}
