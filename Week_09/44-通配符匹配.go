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
	s := "adceb"
	p := "*a*b"
	match := isMatch(s, p)
	fmt.Println(match)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	// 贪心：暂时没写

	// dp优化：一旦 p[j] = '*' 为true，则后面都从 j+1 开始遍历
	m, n := len(s), len(p)
	dp, preCol := make([][]bool, m+1), make([]bool, n+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0], preCol[0] = true, true
	for j := 1; j <= n; j++ {
		dp[0][j] = p[j-1] == '*' && dp[0][j-1]
		preCol[j] = dp[0][j]
	}
	currCol := 1
	for i := 1; i <= m; i++ {
		for j := currCol; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = preCol[j-1]
				if dp[i][j]{
					currCol = j
				}
			}
			if dp[i][j] {
				preCol[j] = true
			}
		}
	}
	return dp[m][n]

	// dp
	//m, n := len(s), len(p)
	//dp, preCol := make([][]bool, m+1), make([]bool, n+1)
	//for i := 0; i <= m; i++ {
	//	dp[i] = make([]bool, n+1)
	//}
	//dp[0][0], preCol[0] = true, true
	//for j := 1; j <= n; j++ {
	//	dp[0][j] = p[j-1] == '*' && dp[0][j-1]
	//	preCol[j] = dp[0][j]
	//}
	//for i := 1; i <= m; i++ {
	//	for j := 1; j <= n; j++ {
	//		if s[i-1] == p[j-1] || p[j-1] == '?' {
	//			dp[i][j] = dp[i-1][j-1]
	//		} else if p[j-1] == '*' {
	//			dp[i][j] = preCol[j-1]
	//		}
	//		if dp[i][j] {
	//			preCol[j] = true
	//		}
	//	}
	//}
	//return dp[m][n]

	// dfs
}

//leetcode submit region end(Prohibit modification and deletion)
