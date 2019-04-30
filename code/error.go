/**
 * name: error类型
 * author: jie
 * note:
 *
 * Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误
 *
 */

package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("出错了！")

	if err != nil {
		fmt.Println(err)
	}

}
