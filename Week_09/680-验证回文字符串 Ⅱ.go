//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。 
//
// 示例 1: 
//
// 
//输入: "aba"
//输出: True
// 
//
// 示例 2: 
//
// 
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
// 
//
// 注意: 
//
// 
// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。 
// 
// Related Topics 字符串
package main

import "fmt"

func main() {
	s := "abca"
	palindrome := validPalindrome(s)
	fmt.Println(palindrome)
}

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	// 个人写法
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return valid(s, i+1, j) || valid(s, i, j-1)
		}
	}
	return true
}

func valid(s string, l int, r int) bool {
	for ; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
