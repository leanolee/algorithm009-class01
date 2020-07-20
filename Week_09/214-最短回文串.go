//给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。 
//
// 示例 1: 
//
// 输入: "aacecaaa"
//输出: "aaacecaaa"
// 
//
// 示例 2: 
//
// 输入: "abcd"
//输出: "dcbabcd" 
// Related Topics 字符串
package main

import "fmt"

func main() {
	s := "aacecaaa" // aacecaaa#aaacecaa：[0 1 0 0 0 1 2 2 0 1 2 2 3 4 5 6 7]
	//s := "abcd" // abcd#dcba
	//"yyabccbayyxxxx"
	palindrome := shortestPalindrome(s)
	fmt.Println(palindrome)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPalindrome(s string) string {
	// KMP
	cache := []uint8(s)
	cache = append(cache, '#')
	for i := len(s) - 1; i >= 0; i-- {
		cache = append(cache, s[i])
	}
	n := len(cache)
	prefix := make([]int, n)
	for i, j := 1, 0; i < n; {
		if cache[i] == cache[j] {
			j++
			prefix[i] = j
			i++
		} else {
			if j > 0 {
				j = prefix[j-1]
			} else {
				i++
			}
		}
	}
	fmt.Println(prefix)
	return string(append(cache[len(s)+1:n-prefix[n-1]], cache[:len(s)]...))

	// 双指针

	// 个人写法：中心法：O(n^2)
	//copyIdx, n := 0, len(s)
	//for mid, l, r := n-1, 0, 0; mid >= 0; mid-- { // 初始中心索引为 n-1
	//	if mid&1 == 0 { // 偶数，指向字符
	//		l, r = mid>>1-1, mid>>1+1
	//	} else { // mid：奇数，指向两字符之间
	//		l, r = mid>>1, mid>>1+1
	//	}
	//	for ; l >= 0; l, r = l-1, r+1 {
	//		if s[l] != s[r] {
	//			break
	//		}
	//	}
	//	if l == -1 {
	//		copyIdx = mid + 1
	//		break
	//	}
	//}
	//ans := make([]uint8, n-copyIdx)
	//for i, j := n-1, 0; i >= copyIdx; i, j = i-1, j+1 {
	//	ans[j] = s[i]
	//}
	//return string(ans) + s
}

//leetcode submit region end(Prohibit modification and deletion)
