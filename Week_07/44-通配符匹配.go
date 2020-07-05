//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。 
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
// 
//
// 两个字符串完全匹配才算匹配成功。 
//
// 说明: 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。 
// 
//
// 示例 1: 
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。 
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
// 
//
// 示例 3: 
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
// 
//
// 示例 4: 
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
// 
//
// 示例 5: 
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false 
// Related Topics 贪心算法 字符串 动态规划 回溯算法
package main

import "fmt"

func main() {
	//s := "adcdb"
	//p := "*a*b"
	s := ""
	p := "*"
	match := isMatch(s, p)
	fmt.Println(match)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	// dp：个人写法，很慢
	m, n := len(s), len(p)
	if m == 0 {
		for _, c := range p {
			if c != '*' {
				return false
			}
		}
		return true
	}
	dp := make([][]bool, m+1)
	dp[0] = make([]bool, n+1)
	dp[0][0] = true
	for i := 0; i < m; i++ {
		dp[i+1] = make([]bool, n+1)
		for j := 0; j < n; j++ {
			switch p[j] {
			case '?':
				dp[i+1][j+1] = dp[i][j]
			case '*':
				// 可以简化
				//for c := 0; c <= i+1; c++ {
				//	if dp[c][j] {
				//		dp[i+1][j+1] = true
				//		break
				//	}
				//}
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
				if dp[i][j] {
					dp[0][j+1] = true
				}
			//case s[i]:
			default:
				if s[i] == p[j] {
					dp[i+1][j+1] = dp[i][j]
				} // else false
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]

	// 递归：比较难判断
	//return isMatchRecursion(s, p, 0, 0)
}

func isMatchRecursion(s string, p string, i int, j int) bool {
	fmt.Println(i, j)
	if i == len(s) {
		if j == len(p) {
			return true
		} else {
			for k := j; k < len(p); k++ {
				if p[j] != '*' {
					return false
				}
			}
			return true
		}
	} else if j == len(p) {
		return i == len(s)
	}
	switch p[j] {
	case '?':
		return isMatchRecursion(s, p, i+1, j+1)
	case '*':
		for k := i; k <= len(s); k++ {
			if isMatchRecursion(s, p, k, j+1) {
				fmt.Println("?", k, j+1)
				return true
			}
		}
	default:
		if s[i] != p[j] {
			return false
		} else {
			return isMatchRecursion(s, p, i+1, j+1)
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
