/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:39:20
 */
package main

import (
	"golang-knowledge/Algorithms/Sort/bubble"
	//"Algorithms/Sort/quick"

	//"golang-knowledge/Algorithms/Sort/shell"
	//"golang-knowledge/Algorithms/Sort/binaryInsert"
	//"golang-knowledge/Algorithms/Sort/insert"

	//"golang-knowledge/Algorithms/Sort/selection"
	//"golang-knowledge/Algorithms/Sort/heap"

	//"golang-knowledge/Algorithms/Sort/merge"

	//"golang-knowledge/Algorithms/Sort/bucket"

	//"golang-knowledge/Algorithms/Sort/radix"

	//"golang-knowledge/Algorithms/Sort/counting"

	"fmt"
	"golang-knowledge/Algorithms/Sort/utils"
	"log"
)

func main() {
	list := utils.GetArrayOfSize(10)
	log.Println("排序前数组=", list)

	// --插入排序--
	//fmt.Println("--插入排序--")
	//insert.Sort(list)

	//fmt.Println("--折半插入排序--")
	//binaryInsert.Sort(list)

	//fmt.Println("--希尔排序--")
	//shell.Sort(list)

	// --交换排序--
	fmt.Println("--冒泡排序--")
	bubble.Sort(list)

	//fmt.Println("--快速排序--")
	//quick.Sort(list)

	// --选择排序--
	//fmt.Println("--直接选择排序--")
	//selection.Sort(list)

	//fmt.Println("--堆排序--")
	//heap.Sort(list)

	//fmt.Println("--归并排序--")
	//merge.Sort(list)

	//fmt.Println("--桶排序--")
	//bucket.Sort(list)

	//fmt.Println("--基数排序--")
	//radix.Sort(list)

	//fmt.Println("--计数排序--")
	//counting.Sort(list)

	log.Println("排序后数组=", list)
}
