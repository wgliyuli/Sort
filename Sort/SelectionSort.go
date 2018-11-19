package Sort

/**
	选择排序
 */
func SelectionSort(a []int, n int){
	if n <= 1 {
		return
	}

	// *利用下标去进行数组值的比较
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		// 数值交换
		a[i], a[min] = a[min], a[i]
	}
	return
}
