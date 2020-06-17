package heap

import (
	"container/heap"
	"fmt"
)

/*
# heap sort 堆排序

//原理：堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法
//内部排序和外部排序：
//稳定性：不稳定
//适用范围：
//算法步骤:
//1.实现堆
//2.利用堆排序
//3.返回排序后的数据
*/


/*功能：
type Interface interface {
    sort.Interface
    Push(x interface{})
    Pop() interface{}
}

这是堆的接口，heap包里面的方法只是提供的一些堆算法操作，要想使用这些算法操作，就必须实现这些接口，每个接口方法都有具体的含义，堆本身的数据结构由这个接口的具体实现决定，可以是数组、列表。
接口方法：
1）sort.Interface
要实现三个接口：
func Len() int，
func Less(i, j int) bool，
func Swap(i, j int)，其中Less方法的实现决定了堆是最大堆还是最小堆。
2）Push(x interface{})
参数列表：x将存到堆中的元素
功能说明：把元素x存放到切片最末尾。
3）Pop() interface{}
返回值：移除切片末尾的那个元素
功能说明：把最后一个元素移除并将其值返回。
*/

// 定义堆
type myHeap []int

// 1、实现排序接口
func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 最小堆的Less方法实现
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// 最大堆的Less方法实现
//func (h *myHeap) Less(i, j int) bool {
//	return (*h)[i] > (*h)[j]
//}

// 2、实现往堆中添加元素
func (h *myHeap)Push(v interface{}) {
	*h = append(*h, v.(int))
}

// 3.实现删除堆中元素
func (h *myHeap)Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

// 定义一个取值方法
func (h *myHeap) Get(n int) int {
	return (*h)[n]
}


// 利用堆进行排序
func Sort(arr []int) []int {
	length := len(arr)

	h := new(myHeap)
	for i:=0; i<length; i++ {
		h.Push(arr[i])
	}

	// 堆排序处理
	heap.Init(h)
	for j:=0; j<h.Len(); j++ {
		arr[j] = h.Get(j)
		fmt.Println(h.Get(j))
	}
	return arr
}

// 问题：打印出的元素被没有全部是有序数组，？？？