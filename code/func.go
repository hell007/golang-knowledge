/**
 * name: func 函数
 * author: jie
 * note:
 *
 * func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
		//这里是处理逻辑代码
		//返回多个值
		return value1, value2
	}
 *
 *
	1:关键字func用来声明一个函数funcName
	2:函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
	3:函数可以返回多个值
	4:上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
	5:如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
	6:如果没有返回值，那么就直接省略最后的返回信息
	7:如果有返回值， 那么必须在函数的外层添加return语句
*
*/

package main

import (
	"fmt"
)

// 1、定义函数
func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}

// 2、多个返回值
func count(a, b int) (sum int, multiplied int) {

	sum = a + b

	multiplied = a * b

	return sum, multiplied
}

// 3、Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参

// arg ...int告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice：

func myfunc(arg ...int) {

	for _, val := range arg {
		fmt.Printf("arg the number is: %d\n", val)
	}
}

// 4、传值与传指针

//当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上.

// 传值
func test(a int) int {
	a = a + 1 //更改 a
	return a
}

// 传指针
func test2(a *int) int {
	*a = *a + 1 //更改 a
	return *a
}

//传指针使得多个函数能操作同一个对象。

//传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。

//Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）

func main() {

	// 1.
	max_xy := max(1, 2)

	fmt.Println("max val=", max_xy)

	// 2.
	s, m := count(3, 4)

	fmt.Println("3 + 4 =", s)

	fmt.Println("3 * 4 =", m)

	// 3.
	myfunc(1, 2, 3, 4, 5)

	// 4.

	x := 4

	fmt.Println("x = ", x) // 应该输出 "x = 4"

	x1 := test(x) //调用test(x)

	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 5"

	fmt.Println("x = ", x) // 应该输出"x = 4"

	y := 4

	y1 := test2(&y) //调用test2(y)

	fmt.Println("y+1 = ", y1) // 应该输出"y+1 = 5"

	fmt.Println("y = ", y) // 应该输出"y = 5"

}
