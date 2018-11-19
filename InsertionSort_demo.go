package main

import (
	"Algorithm/Sort"
	"fmt"
)

func main() {
	var a = []int{4, 5, 6, 1, 3, 2}
	n := len(a)
	Sort.InsertionSort(a, n)
	fmt.Println(a)
}
