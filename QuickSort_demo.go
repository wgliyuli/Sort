package main

import (
	"Algorithm/Sort"
	"fmt"
)

func main() {
	var a = []int{4, 5, 6, 1, 3, 2}
	Sort.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}