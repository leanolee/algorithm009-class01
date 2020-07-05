//给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
// 
//
// 示例: 
//
// 输入: 
//words = ["oath","pea","eat","rain"] and board =
//[
//  ['o','a','a','n'],
//  ['e','t','a','e'],
//  ['i','h','k','r'],
//  ['i','f','l','v']
//]
//
//输出: ["eat","oath"] 
//
// 说明: 
//你可以假设所有输入都由小写字母 a-z 组成。 
//
// 提示: 
//
// 
// 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？ 
// 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何
//实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。 
// 
// Related Topics 字典树 回溯算法
package main

import "fmt"

func main() {
	//words := []string{"oath", "pea", "eat", "rain", "taklfiih"}
	//board := [][]byte{
	//	{'o', 'a', 'a', 'n'},
	//	{'e', 't', 'a', 'e'},
	//	{'i', 'h', 'k', 'r'},
	//	{'i', 'f', 'l', 'v'}}
	words := []string{"a"}
	board := [][]byte{{'a'}, {'a'}}
	strings := findWords(board, words)
	fmt.Println(strings)
}

/*
目标：分析单词搜索 2 用 Tire 树方式实现的时间复杂度，请同学们提交在学习总结中
*/
//leetcode submit region begin(Prohibit modification and deletion)
func findWords(board [][]byte, words []string) []string {
	// Trie：O()
	trie := &Trie{} // 创建 Trie
	for _, word := range words {
		curr := trie
		for _, w := range word {
			idx := w - 'a'
			if curr.next[idx] == nil {
				t := &Trie{}
				curr.next[idx], curr = t, t
			} else {
				curr = curr.next[idx]
			}
		}
		curr.word = word
	}
	ans := make([]string, 0)
	//dx, dy := [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			findWordsDFS(board, &ans, trie, m, n, i, j)
		}
	}
	return ans
}

func findWordsDFS(board [][]byte, ans *[]string, trie *Trie, m int, n int, i int, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '?' {
		return
	}
	c := board[i][j]
	t := trie.next[c-'a']
	if t == nil {
		return
	}
	if t.word != "" {
		*ans, t.word = append(*ans, t.word), ""
		//return
	}
	board[i][j] = '?'
	findWordsDFS(board, ans, t, m, n, i-1, j)
	findWordsDFS(board, ans, t, m, n, i+1, j)
	findWordsDFS(board, ans, t, m, n, i, j-1)
	findWordsDFS(board, ans, t, m, n, i, j+1)
	board[i][j] = c
}

type Trie struct {
	next [26]*Trie
	word string
}

//leetcode submit region end(Prohibit modification and deletion)
