package main

import (
	"Algorithm/Sort"
	"fmt"
)

func main() {
	var a = []int{5, 4, 7, 1, 2, 8, 3, 10, 6, 9}
	n := len(a)
	Sort.BubbleSort(a, n)
	fmt.Println(a)
}
