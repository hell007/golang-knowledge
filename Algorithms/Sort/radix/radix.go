package radix

/*
# radix sort 基数排序

//原理：将整数按位数切割成不同的数字，然后按每个位数分别比较
//内部排序和外部排序：
//稳定性：稳定
//适用范围：待比较的元素必须是“可取基数的”。时间复杂度低（k 因子一般不高），但基数排序的空间复杂度高，内存开销较大。总体上来说，使用范围比较窄
//算法步骤:
1.将所有待比较数值统一为同样的数位长度，数位较短的数前面补零。
2.然后，从最低位开始，依次进行一次排序。这样从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列
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

// 对指定的位进行排序
// bit 可取 1，10，100 等值
func bitSort(arr []int, bit int) []int {
	n := len(arr)
	// 各个位的相同的数统计到 bitCounts[] 中
	bitCounts := make([]int, 10)
	for i := 0; i < n; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1]
	}

	tmp := make([]int, 10)
	for i := n - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num]-1] = arr[i]
		bitCounts[num]--
	}
	for i := 0; i < n; i++ {
		arr[i] = tmp[i]
	}
	return arr
}

func Sort(arr []int) []int {
	max := getMax(arr)
	for bit:=1; max/bit>0; bit*=10 {
		arr = bitSort(arr, bit)
	}
	return arr
}
