/**
 * name: map
 * author: jie
 * note:
 *
 * map也就是Python中字典的概念，它的格式为map[keyType]valueType
 *
 */

package main

import (
	"fmt"
)

func main() {

	// 1、声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int

	// 另一种map的声明方式
	numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	for _, v := range numbers {
		fmt.Println(v)
	}

	fmt.Println(numbers)

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3

	//map过程中需要注意的几点：

	//map默认是无序的，不管是按照 key 还是按照 value 默认都不排序

	//map的长度是不固定的，也就是和slice一样，也是一种引用类型

	//内置的len函数同样适用于map，返回map拥有的key的数量

	//map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11

	//map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

	// 2、map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	val, success := rating["Go"]

	if success {
		fmt.Println("Go=", val)
	} else {
		fmt.Println("Go=", "not found")
	}

	//3、通过delete删除map的元素

	delete(rating, "C") // 删除key为C的元素

	fmt.Println(rating)

	// 4、map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变

	m := make(map[string]string)

	m["key"] = "val"

	n := m

	n["key"] = "value"

	fmt.Println("m=", m)

	fmt.Println("n=", n)

	// 5、make、new操作

	//make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

	// new返回指针

	// make返回初始化后的（非零）值。

	// 6、零值

	// 关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”

	/*
		int     0
		int8    0
		int32   0
		int64   0
		uint    0x0
		rune    0 //rune的实际类型是 int32
		byte    0x0 // byte的实际类型是 uint8
		float32 0 //长度为 4 byte
		float64 0 //长度为 8 byte
		bool    false
		string  ""
	*/

}
