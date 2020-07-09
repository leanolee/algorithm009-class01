//哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’
//t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一
//本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。 
//
//
// 注意：本题相对原题稍作改动，只需返回未识别的字符数 
//
// 
//
// 示例： 
//
// 输入：
//dictionary = ["looked","just","like","her","brother"]
//sentence = "jesslookedjustliketimherbrother"
//输出： 7
//解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
// 
//
// 提示： 
//
// 
// 0 <= len(sentence) <= 1000 
// dictionary中总字符数不超过 150000。 
// 你可以认为dictionary和sentence中只包含小写字母。 
// 
// Related Topics 记忆化 字符串
package main

import (
	"fmt"
	"math"
)

func main() {
	dictionary := []string{"looked", "just", "like", "her", "brother"}
	//sentence := "jesslookedjustliketimherbrother"
	sentence := "jesslookedjust"
	i := respace(dictionary, sentence)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func respace(dictionary []string, sentence string) int {
	// Rabin-Karp + dp：双百
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	base, mod, n := 97, math.MaxInt32, len(sentence)
	dp, memo := make([]int, n+1), make(map[int]bool)
	for _, word := range dictionary {
		memo[rabinKarp(word, base, mod)] = true
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		hash := 0
		for j := i - 1; j >= 0; j-- { // 从后往前找
			hash = (hash*base + int(sentence[j]-'`')) % mod
			if memo[hash] {
				dp[i] = min(dp[i], dp[j])
				if dp[i] == 0 {
					break
				}
			}
		}
	}
	return dp[n]

	// Trie + dp：
	//min := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//root, n := &Trie{}, len(sentence)
	//buildTrie(root, dictionary)
	//dp := make([]int, n+1)
	//for i := 1; i <= n; i++ {
	//	dp[i] = dp[i-1] + 1
	//	curr := root
	//	for j := i - 1; j >= 0; j-- { // 从后往前找
	//		idx := sentence[j] - 'a'
	//		curr = curr.next[idx]
	//		if curr == nil { // 找不到
	//			break
	//		} else if curr.isWord { // 找到了
	//			dp[i] = min(dp[i], dp[j])
	//			if dp[i] == 0 { // 剪枝
	//				break
	//			}
	//		}
	//	}
	//}
	//return dp[n]
}

func rabinKarp(word string, base int, mod int) int {
	hash := 0
	for i := len(word) - 1; i >= 0; i-- {
		hash = (hash*base + int(word[i]-'`')) % mod // '`' = 96
	}
	return hash
}

func buildTrie(root *Trie, arr []string) {
	for _, word := range arr {
		curr := root
		for i := len(word) - 1; i >= 0; i-- { // 重点：必须是倒序
			idx := word[i] - 'a'
			if curr.next[idx] == nil {
				trie := &Trie{}
				curr.next[idx], curr = trie, trie
			} else {
				curr = curr.next[idx]
			}
		}
		curr.isWord = true
	}
}

type Trie struct {
	next   [26]*Trie
	isWord bool
}

//leetcode submit region end(Prohibit modification and deletion)
