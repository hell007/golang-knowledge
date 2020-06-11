/**
 * name: 字符串
 * author: jie
 * note:
 *
 *字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string
 *
 * 字符串是不可变的，但可进行切片操作
 *
 * 声明一个多行的字符串可以通过`来声明：
 */

package main

import "fmt"

func main() {

	// 修改
	s := "cello"

	fmt.Printf("%s\n", s)

	s = "h" + s[1:] // 字符串虽不能更改，但可进行切片操作

	fmt.Printf("%s\n", s)

	// +
	s1 := "hello,"

	s2 := " world"

	str := s1 + s2

	fmt.Printf("%s\n", str)

	// 如果要声明一个多行的字符串可以通过`来声明：

	m := `hello 
		world`

	fmt.Println(m)

}
