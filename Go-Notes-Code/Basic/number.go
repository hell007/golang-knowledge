/**
 * name: 数值类型
 * author: jie
 * note:
 *
 * 整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称
 *
 * 浮点数的类型有float32和float64两种（没有float类型），默认是float64
 *
 * go还支持复数。它的默认类型是complex128（64位实数+64位虚数）。如果需要小一些的，也有complex64(32位实数+32位虚数)。复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位。
 *
 *
 * warn: 需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错
 */

package main

import "fmt"

func main() {

	// 整数
	var a int8 = 1

	var b int8 = 3

	v1 := a + b

	fmt.Println(v1)

	var m int32 = 1

	var n int32 = 2

	v2 := m + n

	fmt.Println(v2)

	// 复数
	var v3 complex64 = 5 + 5i

	fmt.Printf("Value is: %v", v3)

}
