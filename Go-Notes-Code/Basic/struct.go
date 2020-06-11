/**
 * name: struct类型
 * author: jie
 * note:
 *
 */

package main

import (
	"fmt"
)

// 1、struct定义

// 声明一个新的类型
type person struct {
	name string
	age  int
}

// 除了上面这种P的声明使用之外，还有另外几种声明使用方式：

/*1.按照顺序提供初始化值

P := person{"Tom", 25}

2.通过field:value的方式初始化，这样可以任意顺序

P := person{age:24, name:"Tom"}

3.当然也可以通过new函数分配一个指针，此处P的类型为*person

P := new(person)
*/

// 2、struct的匿名字段

// struct定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human // 匿名字段，那么默认Student就包含了Human的所有字段
	sex   string
}

func main() {

	// 1.
	var p person
	p.name = "曹操"
	p.age = 23
	fmt.Printf("我叫%s，今年%d岁了，是该君临天下了吧！", p.name, p.age)

	// 2.

	// 定义一个学生
	joy := Student{Human{"joy", 22, 110}, "男"}
	fmt.Println("姓名：", joy.name)
	fmt.Println("年龄：", joy.age)
	fmt.Println("体重：", joy.weight)
	fmt.Println("性别：", joy.sex)

	// 修改信息
	joy.weight = 120
	fmt.Println("变胖了，现在都：", joy.weight)

	joy.sex = "保密"
	fmt.Println("注重隐私，性别：", joy.sex)

}
