package shell

import "log"

/*
# shell sort 希尔排序

也称递减增量排序算法, 又叫缩小增量排序算法，以科学家 DL．Shell 名字命名。

//原理：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
待整个序列中的记录“基本有序”时，再对全体记录进行依次直接插入排序。
//内部排序和外部排序：内部排序
//稳定性：不稳定
//适用范围: 1万以内用插入排序，以上使用希尔排序
//算法步骤:
//1. 选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
//2. 按增量序列个数 k，对序列进行 k 趟排序；
//3. 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

*/


func insertion(arr []int, start, gap int) {
	length := len(arr)
	log.Printf("gap=%d", gap)

	for i :=  start + gap; i < length; i += gap {
		backup := arr[i]
		j := i - gap
		log.Printf("i=%d", i)
		log.Printf("j=%d", j)

		log.Printf("backup -> arr[%d]:%d, arr[%d]:%d", i, backup, j, arr[j])
		for j >= 0 && backup < arr[j] {
			arr[j + gap] = arr[j]
			log.Printf("大 -> arr[%d]=%d", j + gap, arr[j])
			j -= gap
			log.Printf("j=%d", j)
		}
		arr[j + gap] = backup
		log.Printf("小 -> arr[%d]=%d-----", j + gap, backup)
	}
}

func Sort(arr []int) []int {
	//设定步长增量
	gap := len(arr)/2

	//结束条件，增量为<=0
	for gap > 0 {
		//根据增量，进行分组插入排序
		for pos := 0; pos<gap; pos ++{
			insertion(arr, pos, gap)
		}
		//分组完，组排序后，重新设置增量
		gap /= 2
	}

	return arr
}