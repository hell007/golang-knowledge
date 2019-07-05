package system

type CasbinRule struct {
	Id    int64  `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	PType string `xorm:"varchar(100) index" json:"pType"`
	V0    string `xorm:"varchar(100) index" json:"v0"`
	V1    string `xorm:"varchar(100) index" json:"v1"`
	V2    string `xorm:"varchar(100) index" json:"v2"`
	V3    string `xorm:"varchar(100) index" json:"v3"`
	V4    string `xorm:"varchar(100) index" json:"v4"`
	V5    string `xorm:"varchar(100) index" json:"v5"`
}
