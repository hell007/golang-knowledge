/**
 * name: defer 延迟语句
 * author: jie
 * note:
 *
 * Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题.
 *
 */

package main

import (
	"fmt"
	"os"
)

func ReadWrite() {

	filePath := "D:/Dev/cygwin/work/golang-knowledge/code/city.txt"
	//1、打开文件
	//file, err := os.Open(filePath) // For read access.

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066)

	//2、关闭文件
	defer file.Close()

	if err != nil {
		fmt.Println("打开失败！")
	}

	//3、读取文件
	var b []byte = make([]byte, 4096)

	n, err := file.Read(b)

	if err != nil {
		fmt.Println("Open file Failed", err)
	}

	data := string(b[:n])
	fmt.Println(data)

}

func main() {
	ReadWrite()
}
