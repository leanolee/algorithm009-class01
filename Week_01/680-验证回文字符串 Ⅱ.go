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
	//s := "abca"
	s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	palindrome := validPalindrome(s)
	fmt.Println(palindrome)
}
/*
迭代：O(n),O(1)
	1.通过一个变量 valid 记录出现非回文字符的特殊情况
		lcu、ucul，警惕这中情况下的逻辑判断
	2.valid逻辑：
		valid == 1：第一次出现左右字符不同，先从左边跳过一个字符
			通过 left, right 记录下当前位置
		valid == 2：第二次出现左右字符不同，现在去右边跳过一个字符
			回退到 left, right 记录下的位置
		valid == 3：第三次出现左右字符不同
			说明两边都不行，返回false
*/
//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	// 个人方法
	valid := 0
	left, right := 0, 0
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			valid++
			switch valid {
			case 1:
				left, right = i, j
				i++
			case 2:
				i, j = left, right
				j--
			case 3:
				return false
			}
			continue
		}
		i++
		j--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
