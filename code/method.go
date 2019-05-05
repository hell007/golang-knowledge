/**
 * name: method 面向对象
 * author: jie
 * note:
 *
 * 带有接收者的函数，我们称为method
 *
 * func (r ReceiverType) funcName(parameters) (results)
 *
 * 在使用method的时候重要注意几点:
 *
 * 1.虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
 * 2.method里面可以访问接收者的字段
 * 3.调用method通过.访问，就像struct里面访问字段一样
 */

package main

import (
	"fmt"
	"math"
)

// 1、method使用

// 接收者
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// method
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Ceil(c.radius * c.radius * math.Pi)
}

// 2、method继承

// 接收者
type Human struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Human  // 匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	job   string
}

// Human 定义method
func (h *Human) Introduce() {
	fmt.Printf("你好，我叫%s，今年%d岁。\n", h.name, h.age)
}

// 3、method重写

// Employee的method重写Human的method
func (e *Employee) Introduce() {
	fmt.Printf("你好，我叫%s，今年%d岁，从事%s工作！\n", e.name, e.age, e.job)
}

//Go里面的面向对象是如此的简单，没有任何的私有、公有关键字，通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则.

func main() {

	// 1.

	r := Rectangle{12, 7.8}
	fmt.Println("长方形的面积 = ", r.area())

	c := Circle{3}
	fmt.Println("圆的面积 = ", c.area())

	// 2.

	jay := Student{Human{"jay", 25, "男"}, "北大"}
	jay.Introduce()

	// 3.
	tom := Employee{Human{"tom", 26, "女"}, "工程师"}
	tom.Introduce()

}
