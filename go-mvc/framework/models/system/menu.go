package system

type (
	// 菜单表
	Menu struct {
		Id          int64  `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path        string `xorm:"varchar(64) notnull" json:"path"`
		Url         string `xorm:"varchar(64) notnull" json:"url"`
		Modular     string `xorm:"varchar(64) notnull" json:"modular"`
		Component   string `xorm:"varchar(64) notnull" json:"component"`
		Name        string `xorm:"varchar(64) notnull" json:"name"`
		Icon        string `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive   string `xorm:"varchar(64) notnull" json:"keep_alive"`
		RequireAuth string `xorm:"varchar(64) notnull" json:"require_auth"`
		ParentId    string `xorm:"INT(10) notnull" json:"parent_id"`
		Enabled     string `xorm:"tinyint(1) notnull" json:"enabled"`

		Children []Children `xorm:"-" json:"children"`
	}

	// 子菜单
	Children struct {
		Id2          int64  `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path2        string `xorm:"varchar(64) notnull" json:"path"`
		Modular2     string `xorm:"varchar(64) notnull" json:"modular"`
		Component2   string `xorm:"varchar(64) notnull" json:"component"`
		Name2        string `xorm:"varchar(64) notnull" json:"name"`
		Icon2        string `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive2   string `xorm:"varchar(64) notnull" json:"keep_alive"`
		RequireAuth2 string `xorm:"varchar(64) notnull" json:"require_auth"`
		ParentId2    string `xorm:"INT(10) notnull" json:"parent_id"`
		Enabled2     string `xorm:"tinyint(1) notnull" json:"enabled"`
	}

	// 菜单树
	MenuTreeGroup struct {
		Menu     `xorm:"extends"`
		Children `xorm:"extends"`
	}
)
