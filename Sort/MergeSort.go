package Sort

/**
	归并排序
	步骤：
		1、分解 - 将数组每次折半划分
		2、合并 - 将划分后的数组两两合并后排序
	方法：
		将一组数据，进行左右不断地对半划分（递归）
		然后再对比左右两边的数据再进行合并
 */

// 归并排序
func MergeSort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}

// 归并
func merge(arr []int, low, mid, high int) {
	leftLen := mid - low + 1
	rightLen := high - mid

	arrLeft := make([]int, leftLen+1)
	for i := 0; i < leftLen; i++ {
		arrLeft[i] = arr[low+i]
	}
	arrLeft[leftLen] = 99999 //哨兵牌

	arrRight := make([]int, rightLen+1)
	for j := 0; j < rightLen; j++ {
		arrRight[j] = arr[mid+j+1]
	}
	arrRight[rightLen] = 99999 //哨兵牌

	// 对两个数组进行比较合并
	i, j := 0, 0
	for k := low; k <= high; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}