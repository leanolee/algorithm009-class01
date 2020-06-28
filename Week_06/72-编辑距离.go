//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。 
//
// 你可以对一个单词进行如下三种操作： 
//
// 
// 插入一个字符 
// 删除一个字符 
// 替换一个字符 
// 
//
// 
//
// 示例 1： 
//
// 输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
// 
//
// 示例 2： 
//
// 输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
// 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	//word1 := "horse"
	//word2 := "ros"
	word1 := "sea"
	word2 := "eat"
	distance := minDistance(word1, word2)
	fmt.Println(distance)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	// dp：
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	dp0 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp0[i] = i
	}
	dp[0] = dp0
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i	// 注意是 i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
