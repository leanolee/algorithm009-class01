//给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。 
//
// 一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列
//，而 "AEC" 不是） 
//
// 题目数据保证答案符合 32 位带符号整数范围。 
//
// 
//
// 示例 1： 
//
// 输入：S = "rabbbit", T = "rabbit"
//输出：3
//解释：
//
//如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
//(上箭头符号 ^ 表示选取的字母)
//
//rabbbit
//^^^^ ^^
//rabbbit
//^^ ^^^^
//rabbbit
//^^^ ^^^
// 
//
// 示例 2： 
//
// 输入：S = "babgbag", T = "bag"
//输出：5
//解释：
//
//如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。 
//(上箭头符号 ^ 表示选取的字母)
//
//babgbag
//^^ ^
//babgbag
//^^    ^
//babgbag
//^    ^^
//babgbag
//  ^  ^^
//babgbag
//    ^^^ 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	S := "babgbag"
	T := "bag"
	//S := "rabbbit"
	//T := "rabbit"
	distinct := numDistinct(S, T)
	fmt.Println(distinct)
}

/*
dp：个人写法
	s[i-1] == t[j-1] {
		dp[i][j][j] = dp[i-1][j][j] + dp[i-1][j-1][j-1]
		截止当前字符的个数 + 截止上个字符的个数
	else
		dp[i][j][j] = dp[i-1][j][j]
		当前字符原有的个数
解释：
	1.状态定义：行、列、每一个列截止时的子序列
	2.初始化
		2.1.初始化dp数组
		2.2.初始化dp[i][0][0] = 1
		2.3.拷贝当前行的前一列
	3.根据dp方程计算值
		3.1.s[i-1] != t[j-1]
		3.2.s[i-1] == t[j-1]
			避免：rab rabb，所以要判断左上角

dp：个人写法优化
	s[i-1] == t[j-1]
		dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		原有的个数 + 截止上个字符的个数
	else
		dp[i][j]=dp[i-1][j]
		原有的个数
*/
//leetcode submit region begin(Prohibit modification and deletion)
func numDistinct(s string, t string) int {
	// dp：个人写法再优化
	m, n := len(s), len(t)
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j, last, pre := 0, 0, 1; j < n; j++ {
			last, pre = pre, dp[j]
			if s[i] == t[j] {
				dp[j] += last
			}
		}
	}
	return dp[n-1]

	/*
		[1 1 0 0]
		[1 1 1 0]
		[1 2 1 0]
		[1 2 1 1]
		[1 3 1 1]
		[1 3 4 1]
		[1 3 4 5]
	*/
	// dp：优化二
	//m, n := len(s), len(t)
	//dp, dpPre := make([]int, n+1), make([]int, n+1)
	//dp[0], dpPre[0] = 1, 1
	//for i := 0; i < m; i++ {
	//	for j := 1; j <= n; j++ {
	//		if s[i] == t[j-1] {
	//			dp[j] += dpPre[j-1] // 避免：rab rabb，所以要判断左上角
	//		}
	//	}
	//	copy(dpPre, dp)
	//	fmt.Println(dpPre)
	//}
	//return dpPre[n]

	// dp：个人写法优化
	//m, n := len(s), len(t)
	//dp := make([][]int, m+1)
	//dp[0] = make([]int, n+1)
	//dp[0][0] = 1
	//for i := 1; i <= m; i++ {
	//	dp[i] = make([]int, n+1)
	//	dp[i][0] = 1
	//	for j := 1; j <= n; j++ {
	//		dp[i][j] = dp[i-1][j]
	//		if s[i-1] == t[j-1] {
	//			dp[i][j] += dp[i-1][j-1] // 避免：rab rabb，所以要判断左上角
	//		}
	//	}
	//}
	//fmt.Println(dp)
	//return dp[m][n]

	// dp：个人写法
	//m, n := len(s), len(t)
	//dp := make([][][]int, m+1)	// 1
	//for i := 0; i <= m; i++ {	// 2
	//	dp[i] = make([][]int, n+1)
	//	for j := 0; j <= n; j++ {	// 2.1
	//		dp[i][j] = make([]int, n+1)
	//	}
	//}
	//dp[0][0][0] = 1	// 2.2
	//for i := 1; i <= m; i++ {	// 3
	//	dp[i][0][0] = 1
	//	for j := 1; j <= n; j++ {
	//		copy(dp[i][j], dp[i][j-1]) // 2.3
	//		dp[i][j][j] = dp[i-1][j][j]	// 3.1
	//		if s[i-1] == t[j-1] {	// 3.2
	//			dp[i][j][j] += dp[i-1][j-1][j-1]
	//		}
	//	}
	//}
	//return dp[m][n][n]
}

//leetcode submit region end(Prohibit modification and deletion)
