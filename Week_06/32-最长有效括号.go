//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。 
//
// 示例 1: 
//
// 输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
// 
//
// 示例 2: 
//
// 输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
// 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	//s := ")()())"
	s := "()(())"
	parentheses := longestValidParentheses(s)
	fmt.Println(parentheses)
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	// dp：
	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		switch s[i] {
		case '(':
		case ')':
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
				continue
			}
			if dp[i] > 0 {
				pre := i - dp[i] - 1
				if pre >= 0 && s[i-dp[i]-1] == '(' {
					dp[i+1] = dp[i] + 2
					if pre-1 >= 0 && dp[pre] > 0 {
						dp[i+1] += dp[pre]
					}
				}
			}
		}
	}
	max := 0
	for _, d := range dp {
		if d > max {
			max = d
		}
	}
	return max

	// 个人写法
	//n := len(s)
	//chars := []byte(s)
	//memo := make([]bool, n)
	//for i, j := 0, 0; i < len(chars); {
	//	if chars[i] == ')' && i > 0 && chars[i-1] == '(' {
	//		k := j - 1
	//		for memo[k] {
	//			k--
	//		}
	//		memo[j], memo[k] = true, true
	//		chars = append(chars[:i-1], chars[i+1:]...)
	//		i--
	//		j++
	//	} else {
	//		i++
	//		j++
	//	}
	//}
	//max, count := 0, 0
	//for _, b := range memo {
	//	if b {
	//		count++
	//	} else {
	//		if count > max {
	//			max = count
	//		}
	//		count = 0
	//	}
	//}
	//if count > max {
	//	max = count
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
