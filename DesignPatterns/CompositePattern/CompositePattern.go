package CompositePattern

import "fmt"

/**
意图：将对象组合成树形结构以表示“部分-整体”的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。
主要解决：它在我们树形结构的问题中，模糊了简单元素和复杂元素的概念，客户程序可以向处理简单元素一样来处理复杂元素，从而使得客户程序与复杂元素的内部结构解偶。
何时使用：
表示对象的部分-整体层次结构也就是树形结构。
希望用户忽略组合对象与单个对象的不同，用户将同一地使用组合结构中的所有对象。
如何解决：树枝和叶子实现同一接口，树枝内部组合该接口。
关键代码：树枝内部组合该接口，并且含有内部属性slice，里面放置元素
优点：
高层模块调用简单
节点自由增加
缺点：在使用组合模式时，其叶子和树枝的声明都是实现类，而不是接口，违反了依赖倒置原则。
使用场景：部分、整体场景。
 */

type Employee struct {
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func GetEmployee(name string, dept string, sal int) *Employee {
	employee := new(Employee)
	employee.name = name
	employee.dept = dept
	employee.salary = sal
	return employee
}

func (e *Employee) Add(ee *Employee) {
	e.subordinates = append(e.subordinates, ee)
}

func (e *Employee) Remove(ee *Employee) {
	target := e.subordinates[:0]
	for _, item := range e.subordinates {
		if item == ee {
			target = append(target, item)
		}
	}
	e.subordinates = target
}

func (e *Employee) GetSubordinates() []*Employee {
	return e.subordinates
}

func (e *Employee) ToString() {
	fmt.Println("name:", e.name, "dept:", e.dept, "salary:", e.salary)
}

func (e *Employee) PrintSubordinates() {
	fmt.Println("=============")
	e.ToString()
	for _, headEmployee := range e.GetSubordinates() {
		headEmployee.ToString()
		for _, employee := range headEmployee.GetSubordinates() {
			employee.ToString()
		}
	}
	fmt.Println("=============")
}
