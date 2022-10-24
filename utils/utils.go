package utils

func Swap(arr []int, a, b int) {
	if a != b {
		arr[a], arr[b] = arr[b], arr[a]
	}
}
