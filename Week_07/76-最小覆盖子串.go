//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。 
//
// 示例： 
//
// 输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC" 
//
// 说明： 
//
// 
// 如果 S 中不存这样的子串，则返回空字符串 ""。 
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。 
// 
// Related Topics 哈希表 双指针 字符串 Sliding Window
package main

import "fmt"

func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	//S := "AA"
	//T := "AA"
	window := minWindow(S, T)
	fmt.Println(window)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	// 模板代码：滑动窗体
	min, begin, count := len(s)+1, 0, len(t)
	memo := make(map[uint8]int)
	for _, c := range t {
		memo[uint8(c)]++
	}
	for i, j := 0, 0; j < len(s); j++ {
		// 写法二
		if v, ok := memo[s[j]]; ok {
			if v > 0 { // 关键判断 1
				count--
			}
			memo[s[j]]--
		}
		// 写法一
		//memo[s[j]]--
		//if memo[s[j]] > 0 {
		//	count--
		//}
		for count == 0 {
			curr := j - i + 1
			if curr < min {
				min, begin = curr, i
			}
			if v, ok := memo[s[i]]; ok {
				if v == 0 { // 关键判断 2
					count++
				}
				memo[s[i]]++
			}
			i++
		}
	}
	if min == len(s)+1 {
		return ""
	}
	return s[begin : begin+min]
}

//leetcode submit region end(Prohibit modification and deletion)
