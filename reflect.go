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
	a := Admin{Id: 0, Name: "", Phone: "", Status: 1, Time: 0}
	fmt.Println("====", a)

	//get
	v := reflect.ValueOf(a)
	val := v.FieldByName("Name").String()
	fmt.Println("===", val)

	//set
	b := reflect.ValueOf(&a).Elem()
	b.FieldByName("Id").SetInt(1000)
	b.FieldByName("Name").SetString("曹操")
	b.FieldByName("Phone").SetString("138888888888")
	b.FieldByName("Time").SetInt(time.Now().Unix())
	fmt.Println("====", a)

}
