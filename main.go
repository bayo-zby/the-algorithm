package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	sort.QuickSort2(arr, 0)
	fmt.Println(arr)
}
