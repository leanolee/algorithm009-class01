//给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。 
//
// 
// 如果剩余字符少于 k 个，则将剩余字符全部反转。 
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。 
// 
//
// 
//
// 示例: 
//
// 输入: s = "abcdefg", k = 2
//输出: "bacdfeg"
// 
//
// 
//
// 提示： 
//
// 
// 该字符串只包含小写英文字母。 
// 给定字符串的长度和 k 在 [1, 10000] 范围内。 
// 
// Related Topics 字符串
package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 8
	str := reverseStr(s, k)
	fmt.Println(str)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	// 个人方法
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cache := []byte(s)
	for i := 0; i < len(cache); i += k << 1 {
		for l, r := i, min(i+k-1, len(cache)-1); l < r; l, r = l+1, r-1 {
			cache[l], cache[r] = cache[r], cache[l]
		}
	}
	return string(cache)
}

//leetcode submit region end(Prohibit modification and deletion)
