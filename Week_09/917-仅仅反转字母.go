//给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入："ab-cd"
//输出："dc-ba"
// 
//
// 示例 2： 
//
// 输入："a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
// 
//
// 示例 3： 
//
// 输入："Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
// 
//
// 
//
// 提示： 
//
// 
// S.length <= 100 
// 33 <= S[i].ASCIIcode <= 122 
// S 中不包含 \ or " 
// 
// Related Topics 字符串
package main

import (
	"fmt"
	"unicode"
)

func main() {
	S := "a-bC-dEf-ghIj"
	letters := reverseOnlyLetters(S)
	fmt.Println(letters)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(S string) string {
	// 双指针
	cache := []rune(S)
	for i, j := 0, len(cache)-1; i < j; {
		if !unicode.IsLetter(cache[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(cache[j]) {
			j--
			continue
		}
		cache[i], cache[j] = cache[j], cache[i]
		i++
		j--
	}
	return string(cache)
}

//leetcode submit region end(Prohibit modification and deletion)
