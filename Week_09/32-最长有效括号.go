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
	s := ")(()())"
	parentheses := longestValidParentheses(s)
	fmt.Println(parentheses)
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	// dp：
	//max, n := 0, len(s)
	//dp := make([]int, n+1)
	//for i, c := range s {
	//	if c == ')' {
	//		if i > 0 && s[i-1] == '(' {
	//			dp[i+1] = dp[i-1] + 2
	//		} else if dp[i] > 0 && i-dp[i] > 0 && s[i-1-dp[i]] == '(' {
	//			dp[i+1] = dp[i] + 2 + dp[i-1-dp[i]]
	//		}
	//		if dp[i+1] > max {
	//			max = dp[i+1]
	//		}
	//	}
	//}
	//return max

	// Stack
	max, stack := 0, []int{-1}
	for i, c := range s {
		if c==')'{
			n := len(stack)
			if n > 1 && s[stack[n-1]] == '(' {
				stack = stack[:n-1]
				curr := i - stack[n-2]
				if curr > max {
					max = curr
				}
				continue
			}
		}
		stack = append(stack, i)
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
