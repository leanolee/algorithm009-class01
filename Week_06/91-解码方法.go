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
	//s := "226"
	s := "12343519203510834589" // 6
	//s := "123435"
	decodings := numDecodings(s)
	fmt.Println(decodings)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			if s[i-1] == '0' || s[i-1] > '2' {	// 失败的情况
				return 0
			} else {
				dp[i+1] = dp[i-1]	// 末尾两数是 10 / 20
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {	// 示例：13 / 25
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
