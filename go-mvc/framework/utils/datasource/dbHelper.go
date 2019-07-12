/**
 * name: dbHelper
 * author: jie
 * date: 2019-6-24
 * note: 获取数据库连接
 */
package datasource

import (
	"../../conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// 主库，单例
func MasterEngine() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	master := conf.MasterDbConfig
	// driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	// 	c.User, c.Pwd, c.Host, c.Port, c.DbName)
	// engine, err := xorm.NewEngine(c.DriverName, driveSource)
	engine, err := xorm.NewEngine(master.Dialect, GetConnURL(&master))

	if err != nil {
		golog.Fatalf("dbhelper.DbInstanceMaster, %s", err)
		return nil
	}

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(conf.ShowSQL)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "jie_")
	engine.SetTableMapper(tbMapper)
	engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return engine
}

// 从库，单例
/*func SlaveEngine() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	engine, err := xorm.NewEngine(conf.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err != nil {
		golog.Fatalf("dbhelper.DbInstanceSlave, %s", err)
		return nil
	}

	engine.SetTZLocation(conf.SysTimeLocation)

	slaveEngine = engine
	return engine
}*/

// 获取数据库连接的url
// true：master主库
func GetConnURL(c *conf.DbConf) (url string) {
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.User,
		c.Pwd,
		c.Host,
		c.Port,
		c.DbName,
		c.Charset)
	return
}
