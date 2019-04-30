/**
 * name: if
 * author: jie
 * note:
 */

package main

import (
	"fmt"
)

func main() {

	x := 12

	if x > 10 {
		fmt.Println("x > 10")
	} else {
		fmt.Println("x < 10")
	}

	// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了

	// 计算获取值y,然后根据y返回的大小，判断是否大于10。
	if y := 10; y > 10 {
		fmt.Println("y > 10")
	} else if y < 10 {
		fmt.Println("y < 10")
	} else {
		fmt.Println("y = 10")
	}

	//这个地方如果这样调用就编译出错了，因为y是条件里面的变量
	//fmt.Println(y)

}
