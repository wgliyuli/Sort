package Sort

/**
	快速排序
 */

//func QuickSort(a []int){
//	if len(a) <= 1 {
//		return
//	}
//	quickSort(a, 0, len(a)-1)
//}

func QuickSort(a []int, left, right int) {
	if len(a) <= 1 {
		return
	}

	temp := a[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && a[j] >= temp {
			j--
		}
		if j >= p {
			a[p] = a[j]
			p = j
		}
		if a[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			a[p] = a[i]
			p = i
		}
	}
	a[p] = temp
	if p-left > 1 {
		QuickSort(a, left, p-1)
	}
	if right-p > 1 {
		QuickSort(a, p+1, right)
	}
}
