package bucket

/*
# bucket sort 桶排序
注意:不能进行负数排序

//原理：将两个的有序数列合并成一个有序数列
//内部排序和外部排序：
//稳定性：稳定
//适用范围：
//算法步骤:
1.获取最大值
2.创建一个最大值长度的数组作为桶数组
3.循环待排序数组，将每个数组的值作为桶的键在桶数组上叠加
4.输出数组

这三种排序算法都利用了桶的概念，但对桶的使用方法上有明显差异：

基数排序：根据键值的每位数字来分配桶；
计数排序：每个桶只存储单一键值；
桶排序：每个桶存储一定范围的数值；
*/

func getMax(arr []int) int {
	max := arr[0]
	for i:=0; i<len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func Sort(arr []int) []int {
	length := len(arr)
	if length < 2 { return arr}

	// 1、获取最大值
	max := getMax(arr)

	// 2、新建桶数组
	bucket := make([]int, max+1)
	sortedIndex := 0

	// 3、将待排序数组的值为桶的键，进行统计
	for i:=0; i<length; i++ {
		bucket[arr[i]] += 1
	}

	// 4、将桶的数据输出
	for j:=0; j<=max; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}