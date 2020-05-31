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

import (
	"fmt"
)

func main() {
	sudoku := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(sudoku)
	for _, row := range sudoku {
		for i, n := range row {
			row[i] = n - 48
		}
	}
	for _, row := range sudoku {
		fmt.Println(row)
	}
}

/*
思路：在leetcode-37的基础上，使用回溯法
回溯（数组）：O(9 * 9*8 * 9*8*7 * ... * 9!),O(90 * 3)
回溯（位）：O(9 * 9*8 * 9*8*7 * ... * 9!),O(9 * 3)/O(81)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	// 回溯：使用位运算
	rows, cols, boxes := [9]int{}, [9]int{}, [9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				ok := 1 << (num - 1)
				rows[i] ^= ok
				cols[j] ^= ok
				boxes[i-i%3+j/3] ^= ok
			}
		}
	}
	solve := bitsSolveSudoku(0, 0, rows, cols, boxes, board)	// 现在是，占了的位为 1，没占的位为 0
	fmt.Println(solve)

	// 回溯：使用数组
	//rows := make([][]bool, n) // 9*10
	//cols := make([][]bool, n) // 9*10
	//boxes := make([][]bool, n) // 9*10
	//rows := [9][10]bool{}
	//cols := [9][10]bool{}
	//// row:[0-8],col:[0-8] row-row%3:[0 3 6],col/3:[0 1 2]
	//boxes := [9][10]bool{}
	//for i := 0; i < 9; i++ {
	//	for j := 0; j < 9; j++ {
	//		if board[i][j] != '.' {
	//			num := board[i][j] - '0'
	//			rows[i][num], cols[j][num], boxes[i-i%3+j/3][num] = true, true, true
	//		}
	//	}
	//}
	//solve := solveSudokuRecursion(0, 0, rows, cols, boxes, board)
	//fmt.Println(solve)
}
/*
对比：leetcode-51
	存数据：
		37：已使用：0，未使用：1
		51：已使用：0，未使用：1
	使用数据：
		37：已使用：1，未使用：0
		51：已使用：0，未使用：1
	遍历：
		37：轮询 num，去 位存的数据 中查
		51：轮询 位数据，获取能使用的 num
	回溯：
		37：使用 数组 作为递归参数，需要回溯。因为改变 int 后，还要放进数组中
		51：使用 int 作为递归参数，不需要回溯
*/
func bitsSolveSudoku(r int, c int, rows [9]int, cols [9]int, boxes [9]int, board [][]byte) bool {
	if c == 9 {
		r++
		if r == 9 {
			return true
		}
		c = 0
	}
	if board[r][c] != '.' {	// 正因为必须有这个判断，且只能这样判断，所以需要回溯 board
		return bitsSolveSudoku(r, c+1, rows, cols, boxes, board)
	}
	for num := 1; num < 10; num++ {	// 体会和leetcode-51的区别：这里使用的是遍历数字，51使用的是遍历剩余的可能性
		ok := 1 << (num - 1)
		if rows[r]&ok == ok || cols[c]&ok == ok || boxes[r-r%3+c/3]&ok == ok {
			continue
		}
		board[r][c] = byte(num) + '0'
		rows[r] ^= ok	// ^ 和 | 都可以
		cols[c] ^= ok
		boxes[r-r%3+c/3] ^= ok
		if bitsSolveSudoku(r, c+1, rows, cols, boxes, board) {
			return true
		}
		board[r][c] = '.'
		rows[r] ^= ok	// 为什么这里要回溯，而同为位运算的leetcode-51不用呢？体会使用/不使用回溯的区别
		cols[c] ^= ok	// 因为这里的参数是 数组，leetcode-51的参数是 int
		boxes[r-r%3+c/3] ^= ok	// ^ only
	}
	return false
}

func solveSudokuRecursion(row, col int, rows, cols, boxes [9][10]bool, board [][]byte) bool {
	if col == 9 { // 列超出
		row++
		if row == 9 { // 行超出
			return true
		}
		col = 0
	}
	//for board[row][col] != '.' {	// 最后一个数已经被填，会数组越界
	//	col++
	//}
	if board[row][col] != '.' {
		return solveSudokuRecursion(row, col+1, rows, cols, boxes, board)
	}
	for num := byte(1); num < 10; num++ {
		if rows[row][num] || cols[col][num] || boxes[row-row%3+col/3][num] {
			continue
		}
		rows[row][num], cols[col][num], boxes[row-row%3+col/3][num] = true, true, true
		board[row][col] = num + '0'
		// 必须要有返回值，不然找到答案也会继续回溯
		if solveSudokuRecursion(row, col+1, rows, cols, boxes, board) { // 已经找到最终解，终止所有递归
			return true
		}
		board[row][col] = '.' // 如果 1-9 没有合适的，则回溯
		rows[row][num], cols[col][num], boxes[row-row%3+col/3][num] = false, false, false
	}
	return false // 如果 1-9 都没有合适的，则返回 solve 的默认值 false
}

//leetcode submit region end(Prohibit modification and deletion)
