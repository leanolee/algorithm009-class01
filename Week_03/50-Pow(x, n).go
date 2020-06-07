//实现 pow(x, n) ，即计算 x 的 n 次幂函数。 
//
// 示例 1: 
//
// 输入: 2.00000, 10
//输出: 1024.00000
// 
//
// 示例 2: 
//
// 输入: 2.10000, 3
//输出: 9.26100
// 
//
// 示例 3: 
//
// 输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25 
//
// 说明: 
//
// 
// -100.0 < x < 100.0 
// n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。 
// 
// Related Topics 数学 二分查找
package main

import "fmt"

func main() {
	x := 2.0
	n := -2
	//x := 2.00000
	//n := 10
	pow := myPow(x, n)
	fmt.Println(pow)
}
/*
暴力法：
递归：分治、二分
迭代：分治、二分
*/
//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	// 迭代：2
	N, pow := n, 1.0
	if n < 0 {
		x = 1 / x
		N = -n
	}
	for ; N > 0; N >>= 1 {
		if N&1 == 1 {
			pow *= x
		}
		x *= x
	}
	return pow

	// 迭代：分治
	//if n >= 0 {
	//	return getMyPow(x, n)
	//} else {
	//	return 1 / getMyPow(x, -n)
	//}

	// 递归：分治
	//if n >= 0 {
	//	return myPowRecursion(x, n)
	//} else {
	//	return 1 / myPowRecursion(x, -n)
	//}
}
func getMyPow(x float64, n int) float64 {
	pow := 1.0
	for n > 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
func myPowRecursion(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	pow := 1.0
	if n&1 == 1 {
		pow *= x
	}
	x *= x
	return pow * myPowRecursion(x, n>>1)
}

//leetcode submit region end(Prohibit modification and deletion)
