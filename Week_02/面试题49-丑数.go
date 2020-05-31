//我们把只包含因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。 
//
// 
//
// 示例: 
//
// 输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。 
//
// 说明: 
//
// 
// 1 是丑数。 
// n 不超过1690。 
// 
//
// 注意：本题与主站 264 题相同：https://leetcode-cn.com/problems/ugly-number-ii/ 
// Related Topics 数学
package main

import (
	"./MyTools"
	"fmt"
)

func main() {
	n := 1352
	number := nthUglyNumber(n)
	fmt.Println(number)
}

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	// dp
	dp := make([]int, n)
	dp[0] = 1
	multi_2, multi_3, multi_5 := 0, 0, 0
	for i := 1; i < n; i++ {
		m2, m3, m5 := dp[multi_2]<<1, dp[multi_3]*3, dp[multi_5]*5
		dpI := MyTools.Min(MyTools.Min(m2, m3), m5)
		dp[i] = dpI
		if m2 == dpI {
			multi_2++
		}
		if m3 == dpI {
			multi_3++
		}
		if m5 == dpI {
			multi_5++
		}
	}
	return dp[n-1]

	// 堆？

	// 暴力法
	//count := 0
	//i := 1
	//for ; count < n; i++ {
	//	if isUglyNumber(i) {
	//		count++
	//	}
	//}
	//return i - 1
}
func isUglyNumber(n int) bool {
	for n&1 == 0 {
		n >>= 1
	}
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	for n > 0 && n%5 == 0 {
		n /= 5
	}
	return n == 1
}

//leetcode submit region end(Prohibit modification and deletion)
