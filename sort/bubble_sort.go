package sort

import "algorithm/utils"

// 冒泡排序
func BubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				utils.Swap(arr, j-1, j)
			}
		}
	}
	return arr
}
