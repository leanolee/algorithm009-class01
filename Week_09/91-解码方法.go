//一条包含字母 A-Z 的消息通过以下方式进行了编码： 
//
// 'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
// 
//
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。 
//
// 示例 1: 
//
// 输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// 
//
// 示例 2: 
//
// 输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
// 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	//s := "827478591293823183"
	//s := "226"
	s := "110"
	decodings := numDecodings(s)
	fmt.Println(decodings)
}

/*
dp：
	三种情况：
	s[i] = '0'
		dp[i] = dp[i-2]
	s[i-1] = '1' || s[i-1] = '2' && s[i] <= '6'
		dp[i] = dp[i-1] + dp[i-2]
	else
		dp[i] = dp[i-1]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	// dp
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i, c := range s {
		switch c {
		case '0':
			if i == 0 || s[i-1] > '2' || s[i-1] == '0' {
				return 0
			}
			dp[i+1] = dp[i-1]
		default:
			if i > 0 && (s[i-1] == '2' && c <= '6' || s[i-1] == '1') {
				dp[i+1] = dp[i-1] + dp[i]
			} else {
				dp[i+1] = dp[i]
			}
		}
	}
	return dp[n]

	// 顺推：不通，比较难判断，涉及 *2 *3 等
	//count := 1
	//for i := range s {
	//	switch s[i] {
	//	case '0':
	//		if i == 0 || s[i-1] > '2' || s[i-1] == '0' {
	//			return 0
	//		}
	//	case '1':
	//		if i < len(s)-1 && s[i+1] != '0' {
	//			count <<= 1
	//		}
	//	case '2':
	//		if i < len(s)-1 && s[i+1] != '0' && s[i+1] <= '6' {
	//			count <<= 1
	//		}
	//	}
	//}
	//return count
}

//leetcode submit region end(Prohibit modification and deletion)
