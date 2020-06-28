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
	//s := "ADOBECODEBANC"
	//t := "ABC"
	s := "a"
	t := "aa"
	window := minWindow(s, t)
	fmt.Println("window:", window)
}

/*

 */
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

	// dp：采用 [256]int，参考：

	// dp：滑动窗口，很慢
	m, n, min, left := len(s), len(t), len(s)+1, 0
	memo, window := make(map[uint8]int), make(map[uint8]int) // 不方便用 26，因为有小写
	for _, c := range t {
		memo[uint8(c)]++
	}
	check := func() bool {
		for k, v := range memo {
			if window[k] < v {
				return false
			}
		}
		return true
	}
	for i, j := 0, 0; j < m; j++ {
		window[s[j]]++
		if j < i+n-1 {
			continue
		}
		for check() {	// 此时缩小 i，但是可以只判断移除的字符
			currLen := j - i + 1
			if currLen < min {
				min, left = currLen, i
			}
			window[s[i]]--
			i++
		}
	}
	if min == m+1 {
		return ""
	}
	return s[left : left+min]

	// 暴力：超时
	//min, l, m, n, right := len(s)+1, 0, len(s), len(t), len(s)-len(t)
	//memo := make(map[uint8]int) // map[65:1 66:1 67:1]
	//for _, c := range t {
	//	memo[uint8(c)]++
	//}
	//contain := func(source map[uint8]int) bool {
	//	for k, v := range memo {
	//		if source[k] < v {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//for i := 0; i <= right; i++ { // 固定子串左
	//	source := make(map[uint8]int)
	//	for j := i; j < m; j++ { // 固定子串右
	//		source[s[j]]++
	//		if j < i+n-1 {
	//			continue
	//		}
	//		if contain(source) {
	//			curr := j - i + 1
	//			if curr == n {
	//				return s[i : j+1]
	//			}
	//			if curr < min {
	//				min, l = curr, i
	//			}
	//		}
	//	}
	//}
	//if min == len(s)+1 {
	//	return ""
	//}
	//return s[l : l+min]
}

//leetcode submit region end(Prohibit modification and deletion)
