package sort

import "algorithm/utils"

/*
快速排序
*/
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// 进行原地分组
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l+1] < arr[l] {
			// 当光标位小于后一位时,
			utils.Swap(arr, l, l+1)
			l++
		} else if arr[l+1] > arr[l] {
			utils.Swap(arr, l+1, r)
			r--
		} else {
			l++
		}
	}
	QuickSort(arr[l+1:]) //较小部分排序
	QuickSort(arr[:l])   //较大部分排序
	return arr
}
