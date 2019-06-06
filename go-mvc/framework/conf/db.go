/**
 * name: db
 * author: jie
 * date: 2019-6-4
 * note: 数据库设置
 */

package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "admin",
	DbName: "csms-system",
}

var SlaveDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "admin",
	DbName: "csms-system",
}
