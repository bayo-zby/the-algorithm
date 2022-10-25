package recursion

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + Sum(arr[1:])
}
