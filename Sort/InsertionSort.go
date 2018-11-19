package Sort

/**
	插入排序
 */
func InsertionSort(a []int, n int){
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]	// 数据移动
			} else {
				break
			}
		}
		a[j+1] = value	// 插入数据
	}
	return
}