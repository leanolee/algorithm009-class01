//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。 
//
// 示例 1: 
//
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出: false 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	/*
		[[false false false false false false false]
		[false true false false false false false]
		[false true false false false false false]
		[false true true true true true false]
		[false false true true false true false]
		[false false false true true true true]
		[false false false false true false true]]
	*/
	s3 := "aadbbcbcac"
	//s3 := "aadbbbaccc"
	//s1 := "a"
	//s2 := ""
	//s3 := "a"
	//s1 := "aaaab"
	//s2 := "aaaac"
	//s3 := "aaaaaaaabb"	// 321 次
	interleave := isInterleave(s1, s2, s3)
	fmt.Println(interleave)
}

/*
暴力法：2^z
dfs：
	1.失败补丁
	2.dfs
		2.1.边界条件
		2.2.s1[i] != s2[j]，选择匹配 s1 or s2 作为结果
			ans = dfs(i+1, j, k+1)
			或 ans = dfs(i, j+1, k+1)
			两者只会计算其中之一
		2.3.s1[i] == s2[j]，选择同时匹配 s1 and s2 作为结果
			ans = dfs(i+1, j, k+1) || dfs(i, j+1, k+1)
			两者都要计算
DFS：与上面超时 dfs 唯一不同是：再加一个剪枝
	2.3.s1[i] == s2[j]，选择同时匹配 s1 and s2 作为结果
		if ans = DFS(i+1, j, k+1) = true {
			// 直接返回true，不再计算 DFS(i, j+1, k+1)
		} else {
			// 计算 DFS(i, j+1, k+1)
		}
DP：
	1.预判
	2.状态定义：
		2.1.第0行，第0列方便于 dp方程 的判断
		2.2.补丁 dp[1][1] = true
	3.开始递推
		3.1.i > 1 && s1[i-2] == s3[i+j-3] && dp[i-1][j]
			dp[i][j] = true
		3.2.j > 1 && s2[j-2] == s3[i+j-3] && dp[i][j-1]
			dp[i][j] = true
		3.2.匹配失败
			dp[i][j] = false
*/
//leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {
	// dp：滚动数组
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= m; i++ {
		dp[0] = dp[0] && (i == 0 || s1[i-1] == s3[i-1])
		for j := 1; j <= n; j++ {
			dp[j] = i > 0 && s1[i-1] == s3[i+j-1] && dp[j] || s2[j-1] == s3[i+j-1] && dp[j-1]
		}
	}
	return dp[n]

	// dp：
	//m, n := len(s1), len(s2)
	//if m+n != len(s3) { // 1
	//	return false
	//}
	//dp := make([][]bool, m+2) // 2
	//for i := range dp {
	//	dp[i] = make([]bool, n+2) // 2.1
	//}
	//dp[1][1] = true             // 2.2
	//for i := 1; i <= m+1; i++ { // 3
	//	for j := 1; j <= n+1; j++ {
	//		if i > 1 && s1[i-2] == s3[i+j-3] && dp[i-1][j] || j > 1 && s2[j-2] == s3[i+j-3] && dp[i][j-1] { // 3.1 & 3.2
	//			dp[i][j] = true
	//		} // 3.3
	//	}
	//}
	//return dp[m+1][n+1]

	// DFS：通过
	//if len(s1)+len(s2) != len(s3) {
	//	return false
	//}
	//return DFS(s1, s2, s3, 0, 0, 0)

	// dfs：超时
	//if len(s1)+len(s2) != len(s3) { // 1
	//	return false
	//}
	//return dfs(s1, s2, s3, 0, 0, 0) // 2

	// 暴力

}
func DFS(s1 string, s2 string, s3 string, i int, j int, k int) bool {
	if k == len(s3) {
		return true
	} else if i == len(s1) {
		return s3[k:] == s2[j:]
	} else if j == len(s2) {
		return s3[k:] == s1[i:]
	}
	ans := false
	if s3[k] == s1[i] {
		ans = DFS(s1, s2, s3, i+1, j, k+1)
		if !ans && s1[i] == s2[j] { // 2.3
			ans = DFS(s1, s2, s3, i, j+1, k+1)
		}
	} else if s3[k] == s2[j] {
		ans = DFS(s1, s2, s3, i, j+1, k+1)
	}
	//ans := false	// 这样写更慢
	//if s3[k] == s1[i] {
	//	ans = DFS(s1, s2, s3, i+1, j, k+1)
	//}
	//if !ans && s3[k] == s2[j] { // 2.3
	//	ans = DFS(s1, s2, s3, i, j+1, k+1)
	//}
	return ans
}
func dfs(s1 string, s2 string, s3 string, i int, j int, k int) bool {
	if k == len(s3) { // 2.1
		return true
	} else if i == len(s1) {
		return s3[k:] == s2[j:]
	} else if j == len(s2) {
		return s3[k:] == s1[i:]
	}
	ans := false
	if s3[k] == s1[i] { // 2.2
		ans = dfs(s1, s2, s3, i+1, j, k+1)
	}
	if s3[k] == s2[j] { // 2.2 & 2.3
		ans = dfs(s1, s2, s3, i, j+1, k+1) || ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
