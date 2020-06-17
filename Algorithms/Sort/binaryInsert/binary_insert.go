package binaryInsert

/*
# binary insert sort 折半插入排序

//原理：折半插入排序用到了二分查找法的思想，因此在实现中也要参考二分查找法实现的思想:
将直接插入排序中寻找 A[i] 的插入位置的方法改为采用折半比较，即可得到折半插入排序算法
////内部排序和外部排序：内部排序
//稳定性：稳定
//适用范围：
//算法步骤:
1.计算 0 ~ i-1 的中间点，用 i 索引处的元素与中间值进行比较，如果 i 索引处的元素大，说明要插入的这个元素应该在中间值和刚加入i索引之间，反之，就是在刚开始的位置到中间值的位置，这样很简单的完成了折半；
2.在相应的半个范围里面找插入的位置时，不断的用（1）步骤缩小范围，不停的折半，范围依次缩小为 1/2 1/4 1/8 .......快速的确定出第 i 个元素要插在什么地方；
3.确定位置之后，将整个序列后移，并将元素插入到相应位置。
*/


func Sort(arr []int) []int{
	length := len(arr)
	if length <= 1 { return arr}

	for i := 1; i < length; i++ {
		// 如果待排序值小于前值，说明需要排序
		if arr[i] < arr[i-1] {
			key, low, high := arr[i], 0, i
			for low <= high {
				mid := (low + high) / 2
				if arr[mid] > key {
					high = mid - 1
				} else if arr[mid] < key {
					low = mid + 1
				} else {
					high = mid
					break
				}
			}
			var k int
			// 循环结束后，high 的值便是 key 需要在待排序中的位置
			for k = i - 1; k > high; k-- {
				arr[k+1] = arr[k]
			}
			arr[k+1] = key
		}
	}

	return arr
}