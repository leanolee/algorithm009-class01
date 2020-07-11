//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。 
//
// 示例 1: 
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s = "rat", t = "car"
//输出: false 
//
// 说明: 
//你可以假设字符串只包含小写字母。 
//
// 进阶: 
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？ 
// Related Topics 排序 哈希表
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "anagram"
	t := "nagaram"
	anagram := isAnagram(s, t)
	fmt.Println(anagram)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	// 字符排序
	if len(s) != len(t) {
		return false
	}
	byteS := []byte(s)
	byteT := []byte(t)
	sort.Slice(byteS, func(i, j int) bool {
		return byteS[i] < byteS[j]
	})
	sort.Slice(byteT, func(i, j int) bool {
		return byteT[i] < byteT[j]
	})
	for i := 0; i < len(byteS); i++ {
		if byteS[i] != byteT[i] {
			return false
		}
	}
	return true

	// 26计数
	//if len(s) != len(t) {
	//	return false
	//}
	//chars := [26]int{}
	//for i := 0; i < len(s); i++ {
	//	chars[s[i]-'a']++
	//	chars[t[i]-'a']--
	//}
	//for _, i := range chars {
	//	if i > 0 {
	//		return false
	//	}
	//}
	//return true

	// hash：略
	//if len(s) != len(t) {
	//	return false
	//}
	//cache := make(map[uint8]int)
	//for i, v := range s {
	//	cache[uint8(v)]++
	//	cache[t[i]]--
	//}
	//for _, v := range cache {
	//	if v > 0 {
	//		return false
	//	}
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
