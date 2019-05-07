/**
 * name: 使用go-xorm 操作MySQL数据库
 * author: jie
 * note:
 *
 *
 */

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

//定义结构体(xorm支持双向映射)
type Userinfo struct {
	uid        int       `xorm:"int(10) pk autoincr"` //指定主键并自增
	username   string    `xorm:"varchar(64) unique"`  //唯一的
	department string    `xorm:"varchar(64)"`
	created    time.Time `xorm:"created"`
	//Version    int     `xorm:"version"` //乐观锁
}

//定义orm引擎
var x *xorm.Engine

//创建orm引擎
func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/jie?charset=utf8")

	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(Userinfo)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

//增
func Insert(username string, department string, created time.Time) (string, string, time.Time) {
	u1 := new(Userinfo)
	u1.username = "曹操"
	u1.department = "研发部"
	u1.created = time.Now()

	affected, err := x.Insert(u1)
	if err != nil {
		return affected, false
	}
	return affected, true
}

//删
func Del(uid int) {
	userinfo := new(Userinfo)
	x.Id(uid).Delete(userinfo)
}

//改
func update(uid int, userinfo *Userinfo) bool {
	affected, err := x.ID(uid).Update(userinfo)
	if err != nil {
		log.Fatal("错误:", err)
	}
	if affected == 0 {
		return false
	}
	return true
}

//查
func getinfo(uid int) *Userinfo {
	userinfo := &Userinfo{uid: uid}
	is, _ := x.Get(userinfo)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	fmt.println(userinfo)
	return userinfo
}

func main() {

	x.ShowSQL(true) // 显示SQL的执行, 便于调试分析

	//设置表名下划线转驼峰
	//x.SetMapper(core.SameMapper{})

	init()

	getinfo(8)

}
