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

import (
	"fmt"
	"hash/fnv"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(strs)
	fmt.Println(anagrams)

	new32 := fnv.New32()
	new32.Write([]byte{'a', 'b', 'c'})	// 1134309195
	sum32 := new32.Sum32()
	//new32.Reset()
	//fmt.Println(sum32)
	//new32.Write([]byte{'b', 'a', 'c'})
	//sum32 := new32.Sum32()
	fmt.Println(sum32)	// 513390129
}

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	// hash + 排序：慢
	//memo := make(map[string][]string)
	//for _, str := range strs {
	//	sortStr := bytes(str)
	//	sort.Sort(sortStr)
	//	key := string(sortStr)
	//	memo[key] = append(memo[key], str)
	//}
	//anagrams := make([][]string, 0)
	//for _, v := range memo {
	//	anagrams = append(anagrams, v)
	//}
	//return anagrams

	// hash + 26字符
	memo := [26]int{}
	group := make(map[uint32][]string)
	new32 := fnv.New32()
	for _, str := range strs {
		for _, b := range str {
			memo[b-'a'] = memo[b-'a'] + 1
		}
		hashByte := make([]byte, 0)
		for i, count := range memo {
			if count == 0 {
				continue
			}
			for j := 0; j < count; j++ {
				hashByte = append(hashByte, byte('a'+i))
			}
			memo[i] = 0
		}
		new32.Write(hashByte)
		sum32 := new32.Sum32()
		new32.Reset()
		group[sum32] = append(group[sum32], str)
	}
	anagrams := make([][]string, len(group))
	i := 0
	for _, v := range group {
		anagrams[i] = v
		i++
	}
	return anagrams
}

//leetcode submit region end(Prohibit modification and deletion)
