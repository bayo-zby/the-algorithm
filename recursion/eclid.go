package recursion

/*
欧几里得算法
找到两个数的最大公约数
计算两个非负整数的最大公元数
*/
func Gcd(p, q int) int {
	if q == 0 {
		return p
	}
	return Gcd(q, p%q)
}
