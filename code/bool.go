/**
 * name: 布尔类型 bool
 * author: jie
 */

package main

import "fmt"

var open bool      //全局变量声明 一般声明
var status = false // 忽略类型的声明

func main() {

	success := true // 简短声明

	open = true // 赋值操作

	fmt.Println(open)

	fmt.Println(status)

	fmt.Println(success)

}
