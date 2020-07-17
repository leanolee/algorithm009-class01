//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
// 
//
// 
// 每次转换只能改变一个字母。 
// 转换过程中的中间单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回 0。 
// 所有单词具有相同的长度。 
// 所有单词只由小写字母组成。 
// 字典中不存在重复的单词。 
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。 
// 
//
// 示例 1: 
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。 
// Related Topics 广度优先搜索
package main

import "fmt"

func main() {
	//beginWord := "hit"
	//endWord := "cog"
	//wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "ymain"
	endWord := "oecij"
	wordList := []string{"ymann", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", "ymain"}
	/*
		ymain: "ymann", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", ""
		ymann: "", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", ""
		ymanj: "", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "", "yzcnj", ""
		ymcnj: "", "yycrj", "oecij", "", "yzcrj", "yycij", "xecij", "yecij", "", "yzcnj", ""
		yzcnj: "", "yycrj", "oecij", "", "yzcrj", "yycij", "xecij", "yecij", "", "", ""
		yzcrj: "", "yycrj", "oecij", "", "", "yycij", "xecij", "yecij", "", "", ""
		yycrj: "", "", "oecij", "", "", "yycij", "xecij", "yecij", "", "", ""

		yycij: "", "", "oecij", "", "", "", "xecij", "yecij", "",
		yecij: "", "", "oecij", "", "", "", "xecij", "", "", "", """", ""
		xecij: "", "", "oecij", "", "", "", "", "", "", "", """", ""
	*/
	length := ladderLength(beginWord, endWord, wordList)
	fmt.Println(length)
}

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Trie + 双向BFS：写的乱七八糟
	root := &Trie{}
	insert(root, beginWord)
	for _, word := range wordList { // 构建Trie
		insert(root, word)
	}
	if !search(root, []int32(endWord)) {
		return 0
	}
	begin, end, beginVisited, endVisited := []string{beginWord}, []string{endWord}, map[string]int{beginWord: 1}, map[string]int{endWord: 1} // 重点1：数据结构的定义
	for len(begin) > 0 && len(end) > 0 {
		n := len(begin)
		for i := 0; i < n; i++ {
			chars := []int32(begin[i])
			for j := range chars {
				c := chars[j]
				for k := 'a'; k <= 'z'; k++ {
					if k == c {
						continue
					}
					chars[j] = k
					if beginVisited[string(chars)] > 0 {
						chars[j] = c // 重点3：回溯
						continue
					}
					if search(root, chars) { // 重点4
						if endVisited[string(chars)] > 0 {
							return beginVisited[begin[i]] + endVisited[string(chars)]
						}
						begin, beginVisited[string(chars)] = append(begin, string(chars)), beginVisited[begin[i]]+1
					}
					chars[j] = c
				}
			}
		}
		begin = begin[n:] // 重点2
		if len(begin) > len(end) {
			begin, end, beginVisited, endVisited = end, begin, endVisited, beginVisited
		}
	}
	return 0
}
func insert(root *Trie, word string) {
	for _, w := range word {
		i := w - 'a'
		if root.next[i] == nil {
			trie := &Trie{}
			root, root.next[i] = trie, trie
		} else {
			root = root.next[i]
		}
	}
	root.word = word
}
func search(root *Trie, word []int32) bool {
	for _, w := range word {
		if root = root.next[w-'a']; root == nil {
			return false
		}
	}
	return root.word != ""
}

type Trie struct {
	next [26]*Trie
	word string
}

//leetcode submit region end(Prohibit modification and deletion)
