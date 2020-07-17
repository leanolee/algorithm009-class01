//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。 
//
// 示例 1: 
//
// 
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc" 
// 
//
// 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。 
// Related Topics 字符串
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	words := reverseWords(s)
	fmt.Println(words)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	// 个人写法
	strs := make([]string, 0)
	cache := make([]uint8, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			strs = append(strs, string(cache))
			cache = cache[len(cache):]
		} else {
			cache = append(cache, s[i])
		}
	}
	strs = append(strs, string(cache))
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return strings.Join(strs, " ")
}

//leetcode submit region end(Prohibit modification and deletion)
