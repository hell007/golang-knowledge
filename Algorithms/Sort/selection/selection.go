package selection

/*
# select sort 选择排序

//原理：每一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置
//内部排序和外部排序：
//稳定性：不稳定
//适用范围：数据规模越小越好
//算法步骤:
//1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
//2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
//3.重复第二步，直到所有元素均排序完毕
*/

func Sort(arr []int) []int {
	length := len(arr)

	for i:=0; i<length-1; i++ {
		minIndex := i

		for j:=i+1; j<length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//交换
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}