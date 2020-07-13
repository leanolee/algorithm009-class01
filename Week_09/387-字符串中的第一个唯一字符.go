//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。 
//
// 
//
// 示例： 
//
// s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
// 
//
// 
//
// 提示：你可以假定该字符串只包含小写字母。 
// Related Topics 哈希表 字符串
package main

import "fmt"

func main() {
	s := "leetcode"
	char := firstUniqChar(s)
	fmt.Println(char)
}

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	// 26字符
	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}
	for i, v := range s {
		if c[v-'a'] == 1 {
			return i
		}
	}
	return -1

	// hash
	//cache, count := make(map[int32]int), make([]int, len(s))
	//for idx, c := range s {
	//	i, ok := cache[c]
	//	if !ok {
	//		cache[c] = idx
	//		count[idx]++
	//	} else {
	//		count[i]++
	//	}
	//}
	//for i, v := range count {
	//	if v == 1 {
	//		return i
	//	}
	//}
	//return -1
}

//leetcode submit region end(Prohibit modification and deletion)
