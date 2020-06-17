package merge


/*
# merge sort 归并排序

//原理：将两个的有序数列合并成一个有序数列
//内部排序和外部排序：
//稳定性：稳定
//适用范围：
//算法步骤:
//1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
//2.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
//3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
//4.重复步骤 3 直到某一指针达到序列尾；
//5.将另一序列剩下的所有元素直接复制到合并序列尾。
*/

func Sort(arr []int) []int{
	var s = make([]int, len(arr)/2+1)

	if len(arr) < 2 {return arr}

	mid := len(arr) / 2

	Sort(arr[:mid])
	Sort(arr[mid:])

	if arr[mid-1] <= arr[mid] {
		return arr
	}

	copy(s, arr[:mid])

	l, r := 0, mid

	for i := 0; ; i++ {
		if s[l] <= arr[r] {
			arr[i] = s[l]
			l++

			if l == mid {
				break
			}
		} else {
			arr[i] = arr[r]
			r++
			if r == len(arr) {
				copy(arr[i+1:], s[l:mid])
				break
			}
		}
	}

	return arr
}