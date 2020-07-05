//编写一个程序，通过已填充的空格来解决数独问题。 
//
// 一个数独的解法需遵循如下规则： 
//
// 
// 数字 1-9 在每一行只能出现一次。 
// 数字 1-9 在每一列只能出现一次。 
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。 
// 
//
// 空白格用 '.' 表示。 
//
// 
//
// 一个数独。 
//
// 
//
// 答案被标成红色。 
//
// Note: 
//
// 
// 给定的数独序列只包含数字 1-9 和字符 '.' 。 
// 你可以假设给定的数独只有唯一解。 
// 给定数独永远是 9x9 形式的。 
// 
// Related Topics 哈希表 回溯算法
package main

import "fmt"

func main() {
	//sudoku := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	sudoku := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	solveSudoku(sudoku)
	for _, row := range sudoku {
		for i, n := range row {
			row[i] = n - 48
		}
	}
	for _, row := range sudoku {
		for _, c := range row {
			fmt.Printf("%c  ", c+'0')
		}
		fmt.Println()
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	// 位运算：
	//row, col, box := [9]int{}, [9]int{}, [9]int{}
	//var dfs func(r int, c int) bool
	//dfs = func(r int, c int) bool {
	//	if c == 9 {
	//		r++
	//		if r == 9 {
	//			return true
	//		}
	//		c = 0
	//	}
	//	if board[r][c] != '.' { // 原本已被填的数 || 提前遍历，将数据存到3个数组中
	//		// 写法二
	//		//curr := 1 << (board[r][c] - '1')
	//		//if row[r]&curr != 0 || col[c]&curr != 0 || box[boxIdx]&curr != 0 { //检查数字是否已被前面占领
	//		//	return false
	//		//}
	//		//row[r], col[c], box[boxIdx] = row[r]|curr, col[c]|curr, box[boxIdx]|curr
	//		return dfs(r, c+1)
	//	}
	//	boxIdx := r - r%3 + c/3
	//	for i := 0; i < 9; i++ { // 重点
	//		curr := 1 << i
	//		if row[r]&curr != 0 || col[c]&curr != 0 || box[boxIdx]&curr != 0 {
	//			continue
	//		}
	//		board[r][c], row[r], col[c], box[boxIdx] = byte(i)+'1', row[r]|curr, col[c]|curr, box[boxIdx]|curr
	//		if dfs(r, c+1) { // drill down
	//			return true
	//		}
	//		board[r][c], row[r], col[c], box[boxIdx] = '.', row[r]^curr, col[c]^curr, box[boxIdx]^curr // 回溯
	//	}
	//	return false
	//}
	///*
	//	位运算：两种写法
	//		1.先将已知数添加到数组中，会快非常多
	//		2.DFS时，碰到已知数，才添加到数组中，会超时
	//	重点：
	//		1.for i := 0; i < 9; i++ {
	//			遍历每个位置可能填的数 1-9
	//		2.dfs(0, 0, [9]int{}, [9]int{}, [9]int{})
	//			理解：记录的数组可当成参数，也可以不当成参数
	//*/
	//// 写法一
	//for i := 0; i < 9; i++ { // 先存数据
	//	for j := 0; j < 9; j++ {
	//		if board[i][j] != '.' {
	//			//fmt.Println(i, j)
	//			curr := 1 << (board[i][j] - '1')
	//			row[i], col[j], box[i-i%3+j/3] = row[i]|curr, col[j]|curr, box[i-i%3+j/3]|curr
	//		}
	//	}
	//}
	//dfs(0, 0)
	////dfs(0, 0, [9]int{}, [9]int{}, [9]int{})

	// 数组
	row, col, box := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	getBoxIdx := func(i int, j int) int { return i - i%3 + j/3 }
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				val, idx := board[i][j]-'1', getBoxIdx(i, j)
				row[i][val], col[j][val], box[idx][val] = true, true, true
			}
		}
	}
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if c == 9 {
			r++
			if r == 9 {
				return true
			}
			c = 0
		}
		if board[r][c] != '.' {
			return dfs(r, c+1)
		}
		idx := getBoxIdx(r, c)
		for i := 0; i < 9; i++ {
			if row[r][i] || col[c][i] || box[idx][i] {
				continue
			}
			board[r][c], row[r][i], col[c][i], box[idx][i] = byte(i)+'1', true, true, true
			if dfs(r, c+1) {
				return true
			}
			board[r][c], row[r][i], col[c][i], box[idx][i] = '.', false, false, false
		}
		return false
	}
	dfs(0, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
