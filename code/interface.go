/**
 * name: interface
 * author: jie
 * note:
 *
 * interface是一组method签名的组合，我们通过interface来定义对象的一组行为
 *
 */

package main

import (
	"fmt"
)

// 1、interface类型：interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

// 接收者
type Human struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Human  //匿名字段Human
	school string
}

type Employee struct {
	Human  //匿名字段Human
	job    string
	salary float32
}

// method
func (h Human) Introduce() {
	fmt.Printf("你好，我叫%s，今年%d岁。\n", h.name, h.age)
}

func (h Human) Like(lyrics string) {
	fmt.Println("我喜欢唱歌", lyrics)
	fmt.Printf("\n")
}

func (e Employee) Introduce() {
	fmt.Printf("你好，我叫%s，今年%d岁，从事%s工作！\n", e.name, e.age, e.job)
}

// 定义interface

type Mine interface {
	Introduce()
	Like(lyrics string)
}

// 通过上面的代码，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现

// 2、空interface(interface{})

// 空interface不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。

// 定义a为空接口
var a interface{}

// 3、interface变量存储的类型

// Comma-ok断言

// switch测试

// 4、嵌入interface

func main() {

	// 1.
	jay := Student{Human{"jay", 23, "男"}, "北大"}

	tom := Employee{Human{"Tom", 37, "男"}, "工程师", 5000}

	lily := Student{Human{"lily", 43, "女"}, "北大"}

	//定义Mine类型的变量i
	var i Mine

	//i能存储Student
	i = jay
	i.Introduce()
	i.Like("爱你一万年，期待你的表演！")

	//i也能存储Employee
	i = tom
	i.Introduce()
	i.Like("我是一只小小鸟！")

	//定义了slice Mine
	x := make([]Mine, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = jay, tom, lily

	for _, val := range x {
		val.Introduce()
	}

	// 2.
	fmt.Println("\na可以存储任意类型的数值")

	var b int = 5
	s := "Hello world"
	a = b
	fmt.Println("a的类型是int， a = ", a)

	a = s
	fmt.Println("a的类型是string，a = ", a)

}
