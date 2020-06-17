package main

import (
	//"Algorithms/Sort/bubble"
	//"Algorithms/Sort/quick"

	//"Algorithms/Sort/shell"
	//"Algorithms/Sort/binaryInsert"
	//"Algorithms/Sort/insert"

	//"Algorithms/Sort/selection"
	//"Algorithms/Sort/heap"

	//"Algorithms/Sort/merge"

	"Algorithms/Sort/bucket"
	"Algorithms/Sort/utils"
	"fmt"
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
	//fmt.Println("--冒泡排序--")
	//bubble.Sort(list)

	//fmt.Println("--快速排序--")
	//quick.Sort(list)

	// --选择排序--
	//fmt.Println("--直接选择排序--")
	//selection.Sort(list)

	//fmt.Println("--堆排序--")
	//heap.Sort(list)

	//fmt.Println("--归并排序--")
	//merge.Sort(list)

	fmt.Println("--桶排序--")
	bucket.Sort(list)
	log.Println("排序后数组=", list)
}