/**
 * name: goto
 * author: jie
 * note:
 * 请明智地使用它。用goto跳转到必须在当前函数内定义的标签
 * 会造成内存占用
 */

package main

func main() {

	i := 0

Here: //这行的第一个词，以冒号结束作为标签

	if i < 10 {
		println(i)
		i++
	}

	goto Here //跳转到Here去

}
