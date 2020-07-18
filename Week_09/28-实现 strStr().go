//实现 strStr() 函数。 
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如
//果不存在，则返回 -1。 
//
// 示例 1: 
//
// 输入: haystack = "hello", needle = "ll"
//输出: 2
// 
//
// 示例 2: 
//
// 输入: haystack = "aaaaa", needle = "bba"
//输出: -1
// 
//
// 说明: 
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。 
// Related Topics 双指针 字符串
package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	str := strStr(haystack, needle)
	fmt.Println(str)
}

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	// Boyer-Moore匹配算法

	// Rabin-Karp匹配算法

	// KMP匹配算法

	// Sunday匹配算法
	m, n := len(haystack), len(needle)
	badChar := sundayHelper(needle)
	for i, j := 0, 0; i <= m-n; {
		j = 0
		for j < n && haystack[i+j] == needle[j] {
			j++
		}
		if j == n {
			return i
		}
		if i+n == m {
			break
		}
		i += badChar[haystack[i+n]]
	}
	return -1

	// 后缀蛮力匹配算法
	//m, n := len(haystack), len(needle)
	//for i, j := 0, 0; i <= m-n; i++ {
	//	for j = n - 1; j >= 0 && haystack[i+j] == needle[j]; {
	//		j--
	//	}
	//	if j == -1 {
	//		return i
	//	}
	//}
	//return -1
}

func sundayHelper(p string) [256]int {
	badChar, n := [256]int{}, len(p)
	for i := range badChar {
		badChar[i] = n + 1
	}
	for i, c := range p {
		badChar[c] = n - i
	}
	return badChar
}

//leetcode submit region end(Prohibit modification and deletion)
