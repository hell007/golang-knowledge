/**
 * name: slice "动态数组"
 * author: jie
 * note:
 *
 * var fslice []type
 *
 *
 *
 */

package main

import (
	"fmt"
)

func main() {

	// 1、slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。

	// 声明一个含有10个元素元素类型为int的数组
	var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 声明两个含有int的slice
	var a, b []int

	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]

	// b是数组ar的另一个slice
	b = ar[5:7]
	// b的元素是：ar[5]和ar[6]

	fmt.Println(a)

	fmt.Println(b)

	// 2、slice和array的对应关系图

	//slice的默认开始位置是0，ar[:n]等价于ar[0:n]

	//slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]

	//如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
	//

	// 3、slice有几个有用的内置函数：

	//len 获取slice的长度

	fmt.Printf("长度为= %d \n" , len(a))

	//cap 获取slice的最大容量

	fmt.Printf("最大容量= %d \n", cap(a))

	//append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice

	c := append(a, 6)
	fmt.Println(c)

	//copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	// 复制的元素依次替换 dist[0] dist[1] ... ,并且dist src 的len不变

	dist := a
	src := b

	n := copy(dist, src)
	fmt.Printf("复制了 %d 个元素 \n", n) //复制了多少个元素
	fmt.Println(dist)
	fmt.Println(src)

}
