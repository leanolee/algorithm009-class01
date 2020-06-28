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

import (
	"fmt"
)

func main() {
	words := []string{"oath", "pea", "eat", "rain", "taklfiih"}
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'}}
	strings := findWords(board, words)
	fmt.Println(strings)
}

/*
暴力：遍历每一个word
	O(N*m*m*4^k)：words长度N，矩阵m*m，单词长度k
Trie：O(Nk + m*m*4^k)?
*/
//leetcode submit region begin(Prohibit modification and deletion)
func findWords(board [][]byte, words []string) []string {
	ans := make([]string, 0)
	root := buildTrie(words)
	//order := levelOrder(&root) // 遍历test
	//for _, o := range order {
	//	for _, c := range o {
	//		fmt.Printf("%c ", c)
	//	}
	//	fmt.Println()
	//}
	r, c := len(board), len(board[0])
	var dfs func(ans *[]string, root *WordTrie, i, j int)
	dfs = func(ans *[]string, root *WordTrie, i, j int) {
		if i < 0 || i >= r || j < 0 || j >= c { // 索引错误
			return
		}
		c := board[i][j]
		if c == '?' || root.next[c] == nil { // 字符不存在
			return
		}
		trie := root.next[c] // 下一个 Trie
		if trie.word != "" {
			*ans = append(*ans, trie.word) // 找到一个
			//trie.word = ""                 // 避免重复?
		}
		board[i][j] = '?'
		dfs(ans, trie, i-1, j)
		dfs(ans, trie, i, j-1)
		dfs(ans, trie, i, j+1)
		dfs(ans, trie, i+1, j)
		board[i][j] = c // 回溯
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if root.next[board[i][j]] != nil { // 剪枝
				dfs(&ans, &root, i, j)
			}
		}
	}
	return ans
}

type WordTrie struct {
	next map[byte]*WordTrie
	word string
}

func buildTrie(words []string) WordTrie { // 维护一个 Trie
	root := WordTrie{next: make(map[byte]*WordTrie)}
	for _, w := range words {
		pre := &root // 重点：这里必须是指针，不然下面 pre.word = w 修改不了 word
		for i := 0; i < len(w); i++ {
			if pre.next[w[i]] == nil {
				pre.next[w[i]] = &WordTrie{next: make(map[byte]*WordTrie)}
			}
			pre = pre.next[w[i]]
		}
		pre.word = w // 重点：pre 必须是指针，不然修改不了 word
	}
	return root
}
func levelOrder(root *WordTrie) [][]byte {
	order := make([][]byte, 0)
	stack := []*WordTrie{root}
	for len(stack) > 0 {
		level := make([]byte, 0)
		n := len(stack)
		for i := 0; i < n; i++ {
			m := stack[i].next
			for k, v := range m {
				fmt.Printf("%c %s ", k, v.word)
				level = append(level, k)
				stack = append(stack, v)
			}
			fmt.Println()
		}
		stack = stack[n:]
		order = append(order, level)
	}
	return order
}

//leetcode submit region end(Prohibit modification and deletion)
