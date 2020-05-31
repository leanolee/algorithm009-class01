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

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	anagram := isAnagram(s, t)
	fmt.Println(anagram)
}
/*
hash：O(n),O(26)
数组：O(n),O(n)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	// 数组
	if len(s) != len(t) {
		return false
	}
	memo := [26]uint8{}
	for i := 0; i < len(s); i++ {
		memo[s[i]-'a']++
		memo[t[i]-'a']--
	}
	fmt.Println(memo)
	for _, v := range memo {
		if v != 0 {
			return false
		}
	}
	return true

	// hash
	//if len(s) != len(t) {
	//	return false
	//}
	//memo := make(map[uint8]int)
	//for i := 0; i < len(s); i++ {
	//	memo[s[i]] = memo[s[i]] + 1
	//	memo[t[i]] = memo[t[i]] - 1
	//}
	//for _, v := range memo {
	//	if v != 0 {
	//		return false
	//	}
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
