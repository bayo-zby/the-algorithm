package array

// 一个从小到大的数组中查找一个数的索引
func BinarySearch(arr []int, v int) int {
	if len(arr) < 1 {
		return -1
	}
	for l, r, i := 0, len(arr)-1, len(arr)/2; i >= l && i <= r; {
		if arr[i] == v {
			return i
		} else if arr[i] < v {
			// 索引值大于v
			l = i
		} else {
			r = i
		}
		i = (l + r) / 2
	}
	return -1
}
