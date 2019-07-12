/**
 * name: db
 * author: jie
 * date: 2019-6-22
 * note: 数据库设置
 */

package conf

// 数据库
const (
	ShowSQL     bool   = true
	TablePrefix string = "jie_"
)

type DbConf struct {
	Dialect string
	Host    string
	Port    int
	User    string
	Pwd     string
	DbName  string
	Charset string
}

var MasterDbConfig DbConf = DbConf{
	Dialect: "mysql",
	Host:    "127.0.0.1",
	Port:    3306,
	User:    "root",
	Pwd:     "admin",
	DbName:  "csms-system",
	Charset: "utf8",
}

var SlaveDbConfig DbConf = DbConf{
	Dialect: "mysql",
	Host:    "127.0.0.1",
	Port:    3306,
	User:    "root",
	Pwd:     "admin",
	DbName:  "csms-system",
	Charset: "utf8",
}
