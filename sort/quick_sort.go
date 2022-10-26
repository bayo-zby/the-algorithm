package sort

import "algorithm/utils"

/*
快速排序
*/
func QuickSort(arr []int) []int {
	// 基线条件
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0] // 基准值条件
	same := []int{pivot}
	less := make([]int, 0)
	more := make([]int, 0)
	// 通过与基准值的比较，分为三个区间
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			more = append(more, arr[i])
		} else {
			same = append(same, arr[i])
		}
	}
	return append(append(QuickSort(less), same...), QuickSort(more)...)
}

func QuickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	p := partition(arr, 0)
	QuickSort2(arr[:p])
	QuickSort2(arr[p:])
}

func partition(arr []int, p int) int {
	for i := p; i < len(arr); i++ {
		if arr[i] < arr[p] {
			utils.Swap(arr, i, p)
			p = i
		} else if arr[i] > arr[p] {

		}
	}
	return p
}
