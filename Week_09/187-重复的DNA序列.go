//所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究
//非常有帮助。 
//
// 编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。 
//
// 
//
// 示例： 
//
// 输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//输出：["AAAAACCCCC", "CCCCCAAAAA"] 
// Related Topics 位运算 哈希表
package main

import "fmt"

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	sequences := findRepeatedDnaSequences(s)
	fmt.Println(sequences)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatedDnaSequences(s string) []string {
	// 位运算

	// Rabin-Karp
	hash, base, dna, n := uint32(0), uint32(29), 10, len(s)
	if n < dna {
		return nil
	}
	pow, memo, cache, ans := getPow(base, dna), make(map[uint32]bool), make(map[string]bool), make([]string, 0)
	for i := 0; i < dna; i++ {
		hash = hash*base + uint32(s[i])
	}
	memo[hash] = true
	for i := dna; i < n; i++ {
		hash = hash*base + uint32(s[i])
		hash -= pow * uint32(s[i-dna]) // 后减：getPow(base, dna)；先减：getPow(base, dna-1)
		if memo[hash] {                // 还要再判断字符串是否相等
			cache[s[i-dna+1:i+1]] = true
			//ans = append(ans, s[i-dna+1:i+1])
		} else {
			memo[hash] = true
		}
	}
	for k, _ := range cache {
		ans = append(ans, k)
	}
	return ans
}
func getPow(x uint32, n int) uint32 {
	var ans uint32 = 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
