//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。 
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 
//
// 示例: 
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 
//
// 说明: 
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。 
// Related Topics 字符串 回溯算法
package main

import "fmt"

func main() {
	digits := "23"
	combinations := letterCombinations(digits)
	fmt.Println(combinations)
}

/*
暴力法：
暴力法优化：
	算出有多少个结果，如"29"，count = 3*4 = 12
	不过很麻烦
	参见：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/solution/zhi-jie-ding-wei-zi-fu-wei-zhi-yu-shen-qing-mei-yo/

*/
//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	// DFS：回溯
	combinations := make([]string, 0)
	if len(digits) == 0 {
		return combinations
	}
	letters := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	dfsCombinations(0, nil, &combinations, &letters, digits)	// 可以直接传 nil
	return combinations

	// 递归
	//combinations := make([]string, 0)
	//if len(digits) == 0 {
	//	return []string{}
	//}
	//letters := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	//dfsCombinationsRecursion(digits, &letters, &combinations, nil, 0)
	//return combinations

	// 暴力法：BFS
	//if len(digits) == 0 {
	//	return []string{}
	//}
	//combinations := []string{""}
	//letters := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	//for i := 0; i < len(digits); i++ {
	//	chars := letters[digits[i]-'0'-2]
	//	n := len(combinations)
	//	for j := 0; j < len(chars); j++ {
	//		for k := 0; k < n; k++ {
	//			combinations = append(combinations, combinations[k]+string(chars[j]))
	//		}
	//	}
	//	combinations = combinations[n:]
	//}
	//return combinations
}
func dfsCombinations(idx int, s []byte, combinations *[]string, letters *[][]byte, digits string) {
	if len(s) == len(digits) {
		*combinations = append(*combinations, string(s))
		return
	}
	chars := (*letters)[digits[idx]-'0'-2]
	for i := 0; i < len(chars); i++ {
		s = append(s, chars[i])
		dfsCombinations(idx+1, s, combinations, letters, digits)
		s = s[:len(s)-1]
	}
}
func dfsCombinationsRecursion(digits string, letters *[][]byte, combinations *[]string, s []byte, idx int) {
	if len(s) == len(digits) {
		*combinations = append(*combinations, string(s))
		return
	}
	chars := (*letters)[digits[idx]-'0'-2]
	for i := 0; i < len(chars); i++ {
		//str := append(make([]byte, 0), s...)
		dfsCombinationsRecursion(digits, letters, combinations, append(s, chars[i]), idx+1) // 这里没回溯，都可以的
	}
}

//leetcode submit region end(Prohibit modification and deletion)
