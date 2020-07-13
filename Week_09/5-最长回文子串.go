//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。 
//
// 示例 1： 
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
// 
//
// 示例 2： 
//
// 输入: "cbbd"
//输出: "bb"
// 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	s := "aaaa"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}

/*
类同：lc-647
dp：
	s[i] == s[j]
		j-i==1
			dp[i][j] = 2
		j-i>1 && dp[i+1][j-1]>0
			dp[i][j] = dp[i+1][j-1]+1
	s[i] != s[j]
		dp[i][j] = 0

解释：换个方式，初始化 dp 矩阵记录当前回文串的长度
	1.边界条件
	2.记录max，l
		2.1.更新max，l
		2.3.根据max，l计算返回值
	3.递推
		3.1.i==j
			dp[i][j] = 1
		3.2.s[i] != s[j]
			dp[i][j] = 0
		3.3.s[i] == s[j] && dp[i+1][j-1] > 0
			dp[i][j] = dp[i+1][j-1] + 2
			即 s[i~j] 是回文串
		3.4.s[i] == s[j] && i+1==j
			dp[i][j] = 2
*/
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	// Manacher算法：马拉车

	// 中心扩展算法
	//n := len(s)
	//if n < 2 {
	//	return s
	//}
	//maxLen, l := 1, 0
	//for i := 0; i < n; i++ {
	//	maxLen, l = check(s, i, i+1, maxLen, l)
	//	maxLen, l = check(s, i, i+2, maxLen, l)
	//}
	//return s[l : l+maxLen]

	// dp：lc

	// dp：个人写法
	n := len(s)
	if n < 2 {	// 1
		return s
	}
	dp, max, l := make([][]int, n), 1, 0	// 2
	for i := 0; i < n; i++ {	// 3
		dp[i] = make([]int, n)
		dp[i][i] = 1	// 3.1
		for j := 0; j < i; j++ {
			if s[i] != s[j] {	// 3.2
				continue
			}
			if i-j == 1 {	// 3.4
				dp[j][i] = 2
			} else if dp[j+1][i-1] > 0 {	// 3.3
				dp[j][i] = dp[j+1][i-1] + 2
			}
			if dp[j][i] > max {	// 2.1
				max, l = dp[j][i], j
			}
		}
	}
	return s[l : l+max]	// 2.2

	// 暴力：O(n^3)
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n; j++ {
	//		for l, r := i, j; l < r; l, r = l+1, r-1 {
	//			if s[l] != s[r] {
	//				// 不是回文子串
	//			}
	//		}
	//	}
	//}
}

func check(s string, i int, j int, maxLen int, l int) (int, int) {
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			break
		}
	}
	if j-i-1 > maxLen {
		maxLen, l = j-i-1, i+1
	}
	return maxLen, l
}

//leetcode submit region end(Prohibit modification and deletion)
