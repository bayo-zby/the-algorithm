package sort

import "algorithm/utils"

// ιζ©ζεΊ
func SelectSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		midx := i
		for j := midx; j < len(arr); j++ {
			if arr[j] < arr[midx] {
				midx = j
			}
		}
		utils.Swap(arr, i, midx)
	}
	return arr
}
