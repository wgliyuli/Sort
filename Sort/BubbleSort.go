package Sort

/**
	冒泡排序
	从小到大进行排序
 */
func BubbleSort(a []int, n int) {
	// 当数组长度为1的时候，不作排序处理
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		// 提前退出冒泡循环的标志位
		isChange := false
		// 循环 n-i-1 次
		for j := 0; j < n-1-i; j++ {
			// 如果后一个数比a[j]大，则交换位置
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isChange = true // 表示有数据交换
			}
		}
		// 如果没有数据交换，则提前退出
		if !isChange {
			break
		}
	}
	return
}
