//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。 
//
// 示例 1: 
//
// 
//输入: "abab"
//
//输出: True
//
//解释: 可由子字符串 "ab" 重复两次构成。
// 
//
// 示例 2: 
//
// 
//输入: "aba"
//
//输出: False
// 
//
// 示例 3: 
//
// 
//输入: "abcabcabcabc"
//
//输出: True
//
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
// 
// Related Topics 字符串
package main

import "fmt"

func main() {
	//s := "abcabcabcabc"
	//s := "abaaabaaabaa"
	s := "babbabbabbabbab"
	pattern := repeatedSubstringPattern(s)
	fmt.Println(pattern)
}

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	// KMP
	n := len(s)
	prefix := make([]int, n)
	for i, j := 1, 0; i < n; {
		if s[i] == s[j] {
			j++
			prefix[i] = j
			i++
		} else {
			if j > 0 {
				j = prefix[j-1]
			} else {
				i++
			}
		}
	}
	fmt.Println(prefix)
	return prefix[n-1] > 0 && prefix[n-1]%(n-prefix[n-1]) == 0

	// 个人写法
	//for i, j := 0, 1; j <= len(s)>>1; j++ {
	//	for j <= len(s)>>1 && s[j] != s[i] { // 寻找第一个字符出现的下一个位置
	//		j++
	//	}
	//	if j == len(s) || s[j] != s[i] { // 已经失败
	//		return false
	//	}
	//	for ; i < len(s)-j; i += j { // 匹配全部
	//		if i+j<<1 > len(s) || s[i:i+j] != s[i+j:i+j<<1] {
	//			break
	//		}
	//	}
	//	if i == len(s)-j { // 匹配成功
	//		return true
	//	}
	//	i = 0
	//}
	//return false
}

//leetcode submit region end(Prohibit modification and deletion)
