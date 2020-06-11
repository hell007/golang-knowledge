package main

import (
	"fmt"
	"time"
)

func main() {
	//1. 时间戳---------> 时间
	//获得时间戳
	startTimestamp := time.Now().Unix()

	fmt.Println(startTimestamp)

	//把时间戳转换成时间,并格式化为年月日
	timestr := time.Unix(startTimestamp, 0).Format("2006-01-02 15:04:05")

	fmt.Println(timestr)

	//2. 时间---------> 时间戳
	// 没有带时区信息的时间字符串
	a := "2019-07-30 11:50:00"
	// 解析完得到的时间是一个UTC时间
	aa, err1 := time.Parse("2006-01-02 15:04:05", a)
	// 获取当前的时区信息
	b := time.Now().Location()
	// 将时间解析到当前时区
	bb, err2 := time.ParseInLocation("2006-01-02 15:04:05", a, b)

	fmt.Printf("rawTime:%v, err:%v \n", aa, err1)
	fmt.Printf("locTime:%v, err:%v \n", bb, err2)
}
