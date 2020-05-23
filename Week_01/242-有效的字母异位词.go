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
	s, t := "anagram", "nagaram"
	anagram := isAnagram(s, t)
	fmt.Println(anagram)
}
/*
hash：O(n),O(n)
	1.一次循环：s1 每个字符丢进map中，value为出现的次数
	2.二次循环：s2 同上
	3.三次循环：遍历hash，如果 value != 0，返回false
迭代+26字符：O(n),O(26)
	1.26个字符的数组
	2.遍历两个字符串，数组 对应字符的值 +1/-1
	3.遍历数组，值都为0，返回true
*/
//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	// 26字符法：对应的数组
	if len(s) != len(t) {
		return false
	}
	memo := make([]int, 26)
	for i := 0; i < len(s); i++ {
		memo[s[i]-'a']++
		memo[t[i]-'a']--
	}
	for _, v := range memo {
		if v != 0 {
			return false
		}
	}
	return true

	// 暴力法
	//memo := make(map[int32]int)
	//for _, v := range s {
	//	memo[v]++
	//}
	//for _, v := range t {
	//	memo[v]--
	//}
	//for _, v := range memo {
	//	if v != 0 {
	//		return false
	//	}
	//}
	//return true

	// 超级暴力法：O(m*n)
}

//leetcode submit region end(Prohibit modification and deletion)
