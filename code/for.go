/**
 * name: for
 * author: jie
 * note:
 *for expression1; expression2; expression3 {
		//...
	}
*/

package main

import (
	"fmt"
)

func main() {

	// 1、for

	a := 0

	for i := 0; i < 10; i++ {
		a += i
	}

	fmt.Println("a=", a)

	// 2、while
	//有些时候需要进行多个赋值操作，由于Go里面没有,操作符，那么可以使用平行赋值i, j = i+1, j-1;
	//有些时候如果我们忽略expression1和expression3;
	//其中;也可以省略，那么就变成如下的代码了，这就是while的功能

	b := 1

	for b < 10 {
		b += b
	}

	fmt.Println("b=", b)

	// 3、在循环里面有两个关键操作break和continue	,break操作是跳出当前循环，continue是跳过本次循环。当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置;

	// break和continue还可以跟着标号，用来跳到多重循环中的外层循环

	for j := 10; j > 0; j-- {
		if j == 6 {
			//break
			continue
		}
		fmt.Print(j)

		fmt.Println("")
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1

	// 4、for配合range可以用于读取slice和map的数据

	fslice := make(map[string]string)

	fslice["one"] = "1"

	for k, v := range fslice {
		fmt.Println("fslice's key:", k) //注释会报错 处理如下；
		fmt.Println("fslice's val:", v)
	}

	// 由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值 例如

	for _, v := range fslice {
		fmt.Println("fslice's val:", v)
	}

}
