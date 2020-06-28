//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。 
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。 
//
// 示例 1: 
//
// 
//输入: "abc"
//输出: 3
//解释: 三个回文子串: "a", "b", "c".
// 
//
// 示例 2: 
//
// 
//输入: "aaa"
//输出: 6
//说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
// 
//
// 注意: 
//
// 
// 输入的字符串长度不会超过1000。 
// 
// Related Topics 字符串 动态规划
package main

import "fmt"

func main() {
	s := "aaa"
	substrings := countSubstrings(s)
	fmt.Println(substrings)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	// 马拉车算法：马拉车算法可以在线性时间内找出以任何位置为中心的最大回文串


	// 中心法：在长度为 N 的字符串中，可能的回文串中心位置有 2N-1 个：字母，或两个字母中间
	//count := 0
	//n := len(s)
	//for i := 0; i < (n<<1 - 1); i++ {
	//	for l, r := i>>1, (i+1)>>1; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
	//		count++
	//	}
	//}
	//return count

	// 双指针
	//length := 0
	//n := len(s)
	//for i := 0; i < n-1; i++ {
	//	for j := i + 1; j < n; j++ {
	//		for m, n := i, j; m < n; m, n = m+1, n-1 {
	//			if s[m] != s[n] {
	//				goto out
	//			}
	//		}
	//		length++
	//	out:
	//	}
	//}
	//return length + n
}

//leetcode submit region end(Prohibit modification and deletion)
