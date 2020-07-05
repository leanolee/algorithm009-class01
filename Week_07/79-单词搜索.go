//给定一个二维网格和一个单词，找出该单词是否存在于网格中。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 
//
// 示例: 
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false 
//
// 
//
// 提示： 
//
// 
// board 和 word 中只包含大写和小写英文字母。 
// 1 <= board.length <= 200 
// 1 <= board[i].length <= 200 
// 1 <= word.length <= 10^3 
// 
// Related Topics 数组 回溯算法
package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "ABCB"
	b := exist(board, word)
	fmt.Println(b)
}

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	// dfs
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existDFS(board, word, m, n, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func existDFS(board [][]byte, word string, r, c int, i int, j int, idx int) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= r || j < 0 || j >= c || board[i][j] != word[idx] || board[i][j] == '?' {
		return false
	}
	curr := board[i][j]
	board[i][j] = '?'
	ans := existDFS(board, word, r, c, i-1, j, idx+1) || existDFS(board, word, r, c, i+1, j, idx+1) || existDFS(board, word, r, c, i, j-1, idx+1) || existDFS(board, word, r, c, i, j+1, idx+1)
	board[i][j] = curr
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
