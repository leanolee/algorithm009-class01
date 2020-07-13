//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。 
//
// '.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
// 
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。 
//
// 说明: 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。 
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
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "a*"
//输出: true
//解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
// 
//
// 示例 3: 
//
// 输入:
//s = "ab"
//p = ".*"
//输出: true
//解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
// 
//
// 示例 4: 
//
// 输入:
//s = "aab"
//p = "c*a*b"
//输出: true
//解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
// 
//
// 示例 5: 
//
// 输入:
//s = "mississippi"
//p = "mis*is*p*."
//输出: false 
// Related Topics 字符串 动态规划 回溯算法
package main

import "fmt"

func main() {
	s := "aab"
	p := "c*a*b"
	match := isMatch(s, p)
	fmt.Println(match)
}

/*
https://leetcode-cn.com/problems/regular-expression-matching/solution/ji-yu-guan-fang-ti-jie-gen-xiang-xi-de-jiang-jie-b/
递归：

dp：
	p[j]=='*
		取 0 个 p[j-1]*
			dp[i][j]=dp[i][j-2]
		取 1/多 个 p[j-1]*
			s[i] == p[j-1] || p[j-1] == '.'：dp[i][j]=dp[i-1][j]
	else
		s[i] == p[j] || p[j] == '.'：dp[i][j]=dp[i-1][j-1]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	// dp：
	//m, n := len(s), len(p)
	//dp := make([][]bool, m+1)
	//for i := 0; i <= m; i++ {
	//	dp[i] = make([]bool, n+1)
	//}
	//dp[0][0] = true
	//for i := 0; i <= m; i++ {
	//	for j := 1; j <= n; j++ {
	//		if p[j-1] == '*' {
	//			dp[i][j] = dp[i][j-2]                           // 取0个 p[j-2]*
	//			if i > 0 && s[i-1] == p[j-2] || p[j-2] == '.' { // 取1个或多个 p[j-2]*
	//				dp[i][j] = dp[i][j] || dp[i-1][j] // 不是 dp[i-1][j-2]
	//			}
	//		} else if i > 0 && s[i-1] == p[j-1] || p[j-1] == '.' {
	//			dp[i][j] = dp[i-1][j-1]
	//		}
	//	}
	//}
	//return dp[m][n]

	// 递归
	return isMatchRecursion(s, p, 0, 0)
}

func isMatchRecursion(s string, p string, i int, j int) bool {
	if j == len(p) { // i == m 不一定返回false，如：s := "aa" p := "a*"
		return i == len(s)
	}
	firstMatch := i < len(s) && s[i] == p[j] || p[j] == '.'
	if j < len(p)-1 && p[j+1] == '*' {	// 0个 或 多个
		return isMatchRecursion(s, p, i, j+2) || firstMatch && isMatchRecursion(s, p, i+1, j)
	}
	return firstMatch && isMatchRecursion(s, p, i+1, j+1)
}

//leetcode submit region end(Prohibit modification and deletion)
