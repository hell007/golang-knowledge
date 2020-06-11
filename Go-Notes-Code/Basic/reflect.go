/**
 * name: reflect
 * author: jie
 * note:
 *
 * Go语言实现了反射，所谓反射就是能检查程序在运行时的状态
 *
 	 t := reflect.TypeOf(i)
 	 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	 v := reflect.ValueOf(i) 得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
 *
*/

package main

import (
	"fmt"
	"reflect"
	"time"
)

type Admin struct {
	Id     int
	Name   string
	Phone  string
	Status int
	Time   int64
}

func main() {

	//初始化
	a := Admin{Id: 0, Name: "小曹", Phone: "13888888887", Status: 1, Time: 0}
	fmt.Println("====", a)

	//get
	v := reflect.ValueOf(a)
	name := v.FieldByName("Name").String()
	fmt.Println("===", name)

	//set
	b := reflect.ValueOf(&a).Elem()
	b.FieldByName("Id").SetInt(1000)
	b.FieldByName("Name").SetString("曹操")
	b.FieldByName("Phone").SetString("138888888888")
	b.FieldByName("Time").SetInt(time.Now().Unix())
	fmt.Println("====", a)

	//save
	admin := Admin{}
	admin.Id = 2000

	fmt.Println("admin===", admin)

}
