//给定一个字符串，逐个翻转字符串中的每个单词。 
//
// 
//
// 示例 1： 
//
// 输入: "the sky is blue"
//输出: "blue is sky the"
// 
//
// 示例 2： 
//
// 输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 
//
// 示例 3： 
//
// 输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
// 
//
// 
//
// 说明： 
//
// 
// 无空格字符构成一个单词。 
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。 
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。 
// 
//
// 
//
// 进阶： 
//
// 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。 
// Related Topics 字符串
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a good   example"
	words := reverseWords(s)
	fmt.Println(words)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	// 个人写法二
	if len(s) == 0 {
		return s
	}
	l, r := 0, len(s)-1
	for ; l <= r; l++ { // 去头
		if s[l] != 32 {
			break
		}
	}
	strs, i := make([]string, 0), l
	for ; l <= r; l++ { // 分割
		if s[l] == 32 && s[l-1] != 32 {
			strs = append(strs, s[i:l])
			for l <= r && s[l] == 32 {
				l++
			}
			i = l
		}
	}
	if s[r] != 32 {
		strs = append(strs, s[i:l])
	}
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 { // 反转
		strs[i], strs[j] = strs[j], strs[i]
	}
	join := strings.Join(strs, " ") // 拼接
	return join

	// 个人写法
	cache, e := make([]uint8, 0), 0
	for e < len(s) && s[e] == 32 {
		e++
	}
	for i := len(s) - 1; i >= e; i-- {
		if s[i] == 32 && (i == len(s)-1 || s[i+1] == 32) {
			continue
		}
		cache = append(cache, s[i])
	}
	for i, j := 0, 0; i < len(cache); i++ {
		if cache[i] != 32 {
			j = i
			for i < len(cache) && cache[i] != 32 {
				i++
			}
			for k := i - 1; j < k; j, k = j+1, k-1 {
				cache[j], cache[k] = cache[k], cache[j]
			}
		}
	}
	return string(cache)
}

//leetcode submit region end(Prohibit modification and deletion)
