package sort

import "algorithm/utils"

// 插入排序
func InsertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			utils.Swap(arr, j, j+1)
		}
	}
	return arr
}
