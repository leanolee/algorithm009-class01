//给定一个正整数 n，返回长度为 n 的所有可被视为可奖励的出勤记录的数量。 答案可能非常大，你只需返回结果mod 109 + 7的值。 
//
// 学生出勤记录是只包含以下三个字符的字符串： 
//
// 
// 'A' : Absent，缺勤 
// 'L' : Late，迟到 
// 'P' : Present，到场 
// 
//
// 如果记录不包含多于一个'A'（缺勤）或超过两个连续的'L'（迟到），则该记录被视为可奖励的。 
//
// 示例 1: 
//
// 
//输入: n = 2
//输出: 8 
//解释：
//有8个长度为2的记录将被视为可奖励：
//"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//只有"AA"不会被视为可奖励，因为缺勤次数超过一次。 
//
// 注意：n 的值不会超过100000。 
// Related Topics 动态规划
package main

import (
	"fmt"
)

func main() {
	n := 300
	record := checkRecord(n)
	fmt.Println(record)
	//for i := 1; i <= 36; i++ {
	//	record := checkRecord(i)
	//	fmt.Println(record)
	//}
	tar := 82876089                // -168993599
	fmt.Println(tar*2 - 334745777) // 165752178
}

/*
可被奖励 = 所有排列情况 - 不被奖励的情况
*/
//leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(n int) int {
	// dp：分别记录情况
	const M = 1000000007
	a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 := 1, 0, 0, 0, 0, 0
	for i := 1; i <= n; i++ {
		//nextA0l0 := a0l0 + a0l1 + a0l2
		//a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 = nextA0l0, a0l0, a0l1, nextA0l0+a1l0+a1l1+a1l2, a1l0, a1l1
		a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 = (a0l0+a0l1+a0l2)%M, a0l0, a0l1, (a0l0+a0l1+a0l2+a1l0+a1l1+a1l2)%M, a1l0, a1l1
	}
	return (a0l0 + a0l1 + a0l2 + a1l0 + a1l1 + a1l2) % M

	// 递推公式 + dp
	//const M = 1000000007
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			dp[i] = 1 // 必须是 1
		case 1:
			dp[i] = 2
		case 2:
			dp[i] = 4
		case 3:
			dp[i] = 7
		default:
			dp[i] = (dp[i-1]<<1 + M - dp[i-4]) % M
		}
	}
	fmt.Println(dp)
	res := dp[n]
	for i := 1; i <= n; i++ {
		res += (dp[i-1] * dp[n-i]) % M
	}
	return res % M

	// 暴力：超时
	return checkRecordRecursion(n, 0, false, 0)
}

func checkRecordRecursion(n int, curr int, a bool, l int) int {
	if curr == n {
		return 1
	}
	res := 0
	if !a {
		res += checkRecordRecursion(n, curr+1, true, 0)
	}
	if l < 2 {
		res += checkRecordRecursion(n, curr+1, a, l+1)
	}
	return res + checkRecordRecursion(n, curr+1, a, 0)%1000000007
}

//leetcode submit region end(Prohibit modification and deletion)
