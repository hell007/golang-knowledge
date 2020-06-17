package quick

/*
# quick sort 快速排序

//原理：快速排序应该算是在冒泡排序基础上的递归分治法。之所以快，跳跃式的交换消除逆序。
//内部排序和外部排序：
//稳定性：稳定
//适用范围：
//算法步骤:
//1.从数列中挑出一个元素，称为 "基准"（pivot）;
//2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
//4.递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。虽然一直递归下去，但是这个算法总会退出，因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。
*/

//分区操作
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		//这里是关键，找到一个比基准大的数和一个比基准小的数进行交换
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	//将基准交换到中间位置
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right) //找分区点
		//递归地把小于基准值元素的子数列和大于基准值元素的子数列排序
		quickSort(arr, left, partitionIndex-1)
		quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func Sort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}


