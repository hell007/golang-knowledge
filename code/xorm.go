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
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

//定义结构体(xorm支持双向映射)
type Userinfo struct {
	Uid        int       `xorm:"int(10) pk autoincr"` //指定主键并自增
	Username   string    `xorm:"varchar(64) unique"`  //唯一的
	Department string    `xorm:"varchar(64)"`
	Created    time.Time `xorm:"created"`
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
func Insert(userinfo *Userinfo) {
	affected, err := x.Insert(userinfo)
	if err != nil {
		log.Fatal("添加失败:", err)
	}
	fmt.Println(affected)
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
	userinfo := &Userinfo{Uid: uid}
	is, _ := x.Get(userinfo)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	return userinfo
}

func main() {

	x.ShowSQL(true) // 显示SQL的执行, 便于调试分析

	x.SetMapper(core.SameMapper{}) // 设置表名下划线转驼峰

	x.Sync2(new(Userinfo))

	// 1.查询
	fmt.Println("\n struct查询")
	info := getinfo(8)
	fmt.Println("select = ", info, "\n")

	//var username string
	id := 8
	is1, _ := x.Where("uid = ?", 8).Cols("username").Get(&id)
	fmt.Println("查询成功 is1= ", is1, "\n")

	//
	uid := 8
	is2, _ := x.SQL("select username from userinfo").Get(&uid)
	fmt.Println("查询成功 is2=", is2, "\n")

	data1, _ := x.Query("select * from userinfo")
	fmt.Println(" data1 = ", data1, "\n")

	data2, _ := x.QueryString("select * from userinfo")
	fmt.Println(" data2 = ", data2, "\n")

	// 2.添加
	fmt.Println("\n添加")

	// u1 := new(Userinfo)
	// u1.Username = "曹丕"
	// u1.Department = "财务部"
	// u1.Created = time.Now()
	//Insert(u1)

	// 3.更改
	fmt.Println("\n更改")

	// u4 := new(Userinfo)
	// u4.Username = "刘小备1"
	// u4.Department = "刘家桥"
	// update(9, u4)

	// 4.删除
	fmt.Println("\n删除")

	// Del(10)

}
