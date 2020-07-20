//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。 
//
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。 
//
// 说明： 
//
// 
// 字母异位词指字母相同，但排列不同的字符串。 
// 不考虑答案输出的顺序。 
// 
//
// 示例 1: 
//
// 
//输入:
//s: "cbaebabacd" p: "abc"
//
//输出:
//[0, 6]
//
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
// 
//
// 示例 2: 
//
// 
//输入:
//s: "abab" p: "ab"
//
//输出:
//[0, 1, 2]
//
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
// 
// Related Topics 哈希表
package main

import "fmt"

func main() {
	//s := "cbaebabacd"
	//p := "abc"
	s := "abab"
	p := "ab"
	anagrams := findAnagrams(s, p)
	fmt.Println(anagrams)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	// [26]字符：双指针
	ans, cache, curr, n := make([]int, 0), [26]int{}, [26]int{}, len(p)-1
	for _, c := range p {
		cache[c-'a']++
	}
	for l, r := 0, 0; r < len(s); r++ {
		idx := s[r] - 'a'
		curr[idx]++
		for ; curr[idx] > cache[idx]; l++ {
			curr[s[l]-'a']--
		}
		if r-l == n {
			ans = append(ans, l)
		}
	}
	return ans

	// [26]字符：队列
	//ans, cache, curr, count, n := make([]int, 0), [26]int{}, [26]int{}, 26, len(p)
	//for _, c := range p {
	//	idx := c - 'a'
	//	if cache[idx] == 0 {
	//		count--
	//	}
	//	cache[idx]++
	//}
	//for i, c := range s {
	//	in := c - 'a'
	//	if curr[in] == cache[in] {
	//		count--
	//	} else if curr[in] == cache[in]-1 {
	//		count++
	//	}
	//	curr[in]++
	//	if i >= n {
	//		out := s[i-n] - 'a'
	//		if curr[out] == cache[out] {
	//			count--
	//		} else if curr[out] == cache[out]+1 {
	//			count++
	//		}
	//		curr[out]--
	//	}
	//	if count == 26 {
	//		ans = append(ans, i-n+1)
	//	}
	//	//fmt.Printf("%c %d\n", c, count)
	//}
	//return ans

	// hash + 排序
}

//leetcode submit region end(Prohibit modification and deletion)
