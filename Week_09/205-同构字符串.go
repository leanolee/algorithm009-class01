//给定两个字符串 s 和 t，判断它们是否是同构的。 
//
// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。 
//
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。 
//
// 示例 1: 
//
// 输入: s = "egg", t = "add"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s = "foo", t = "bar"
//输出: false 
//
// 示例 3: 
//
// 输入: s = "paper", t = "title"
//输出: true 
//
// 说明: 
//你可以假设 s 和 t 具有相同的长度。 
// Related Topics 哈希表
package main

import "fmt"

func main() {
	s := "aba"
	t := "baa"
	isomorphic := isIsomorphic(s, t)
	fmt.Println(isomorphic)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	// lc
	sPath, tPath := helper(s), helper(t)
	for i, v := range sPath {
		if v != tPath[i] {
			return false
		}
	}
	return true

	// 个人写法
	//memoS := make(map[uint8]uint8)
	//memoT := make(map[uint8]uint8)
	//for i := 0; i < len(s); i++ {
	//	v, ok := memoS[s[i]]
	//	if !ok {
	//		if _, ok := memoT[t[i]]; ok {
	//			return false
	//		}
	//		memoS[s[i]] = t[i]
	//		memoT[t[i]] = s[i]
	//	} else if v != t[i] {
	//		return false
	//	}
	//}
	//for k, v := range memoS {
	//	if memoT[v] != k {
	//		return false
	//	}
	//}
	//return true
}

func helper(s string) []int {
	path := make([]int, len(s))
	memo, count := make(map[int32]int), 1
	for i, v := range s {
		if memo[v] == 0 {
			memo[v] = count
			count++
		}
		path[i] = memo[v]
	}
	return path
}

//leetcode submit region end(Prohibit modification and deletion)
