//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。 
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 
// 
// Related Topics 哈希表 字符串
package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(strs)
	fmt.Println(anagrams)
}

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	// hash：O(n * (k + 26))
	cache, memo, ans := [26]int{}, make(map[string]*[]string), make([][]string, 0)
	for _, s := range strs {
		for _, c := range s {
			cache[c-'a']++
		}
		key, i := make([]uint8, len(s)), 0
		for j, n := range cache {
			for ; n > 0; n-- {
				key[i], i = uint8(j), i+1
			}
			cache[j] = 0
		}
		k := string(key)
		if memo[k] == nil {
			memo[k] = new([]string)
		}
		*memo[k] = append(*memo[k], s)
	}
	for _, v := range memo {
		ans = append(ans, *v)
	}
	return ans

	// 排序：O(n * k log k)

}

//leetcode submit region end(Prohibit modification and deletion)
