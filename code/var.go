/**
 * name: 变量
 * author: jie
 */

package main

import "fmt"

func main() {

	//定义单个变量
	//var v type = val

	var v1 string = "变量1"

	var v2 = "变量2"

	fmt.Println(v1)
	fmt.Println(v2)

	//定义多个变量

	var v3, v4 string = "v3", "v4"
	//var v3, v4 = "v3", "v4"

	fmt.Println(v3)
	fmt.Println(v4)

	//变量简写方式

	v5, v6 := "v5", "v6"

	fmt.Println(v5)
	fmt.Println(v6)

}
