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

import (
	"fmt"
)

func main() {
	x := 2.0
	n := 10
	//x := 2.0034000
	//n := -2
	//x := 0.00001
	//n := 2147483647
	//x := 2.00000
	//n := -2147483648
	pow := myPow(x, n)
	fmt.Println(pow)
}

/*
关键：
	n=37：100101 = 2^5 + 2^2 + 2^0 = 32 + 4 + 1
		powPow = x 的 2^k，即 32 + 4 + 1
		每次 n>>1，则 pow 每遇到奇数就 乘一次 powPow
		最后 pow 乘了，32 + 4 + 1 个 x
*/
//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	// 迭代
	//if n >= 0 {
	//	return quickMulti(x, n)
	//} else {
	//	return 1.0 / quickMulti(x, -n)
	//}

	// 递归
	if n >= 0 {
		return quickMultiRecursion(1.0, x, n)
	} else {
		return 1.0 / quickMultiRecursion(1.0, x, -n)
	}

	// 个人方法：挫
	//pow := x
	//if n == 0 {
	//	return 1
	//}
	//if pow == 1.0 {
	//	return pow
	//}
	//if pow == -1.0 {
	//	if n < 0 {
	//		return -pow
	//	}
	//	return pow
	//}
	//i := n
	//if n < 0 {
	//	i = -i
	//}
	//for ; i > 1; i-- {
	//	pow *= x
	//	if pow == 0.0 {
	//		break
	//	}
	//	if pow >= math.MaxFloat64 {
	//		break
	//	}
	//}
	//if n < 0 {
	//	pow = 1.0 / pow
	//}
	//return pow
}

// 无返回值
func quickMultiRecursion(pow float64, x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	res := 1.0
	if n&1 == 1 {
		res = x
	}
	res *= quickMultiRecursion(pow, x*x, n>>1)
	return res
}

// 有返回值
//func quickMultiRecursion(pow float64, x float64, n int) float64 {
//	if n == 0 {
//		return pow
//	}
//	if n&1 == 1 {
//		pow *= x
//	}
//	pow = quickMultiRecursion(pow, x*x, n>>1)
//	return pow
//}

func quickMulti(x float64, n int) float64 {
	pow := 1.0
	powPow := x
	for n > 0 {
		if n&1 == 1 {
			pow *= powPow
		}
		powPow *= powPow
		n >>= 1
	}
	return pow
}

//leetcode submit region end(Prohibit modification and deletion)
