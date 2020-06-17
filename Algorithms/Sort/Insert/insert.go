package Insert

import "log"

/*
# insert sort 插入排序

插入排序可以原地(in-place)也可以非原地(not-in-place)，
非原地方式是申请新的数组然后把被排数插入新数组，原地方式是备份被排数然后从头部（或尾部）比较、移动来建立有序数组

//原理：通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
////内部排序和外部排序：内部排序
//稳定性：稳定
//适用范围：少量数据的排序
//算法步骤:
// 1.第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列
// 2.从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面

*/

func Sort(arr []int) []int {
	length := len(arr)
	if length <= 1 { return arr}

	for i:=1; i<length; i++ {
		backup := arr[i]
		log.Printf("No%d次比对，backup==%d", i, backup)
		j := i-1;
		log.Println("j==", j)

		// 注意j >= 0必须在前边，否则会数组越界
		for j>=0 && backup < arr[j] {
			// 移动有序数组
			arr[j+1] = arr[j]
			log.Printf("arr[%d]=%d", j+1, arr[j])
			// 反向移动下标
			j--
		}
		// 插队插入移动后的空位
		arr[j+1] = backup
		log.Printf("arr[%d]=%d", j+1, backup)
	}
	return arr
}