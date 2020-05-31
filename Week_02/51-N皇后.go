//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。 
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
//
// 示例: 
//
// 输入: 4
//输出: [
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
// 
//
// 
//
// 提示： 
//
// 
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步
//，可进可退。（引用自 百度百科 - 皇后 ） 
// 
// Related Topics 回溯算法
package main

import (
	"fmt"
)

func main() {
	n := 4
	queens := solveNQueens(n)
	fmt.Println(queens)
}
/*
回溯：
位运算：
*/
//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	// 位运算
	queens := make([][]string, 0)
	solve := make([]int, n)
	dfsBits(0, 0, 0, 0, solve, &queens, n)
	//dfsBits(^0, ^0, ^0, 0, solve, &queens, n)
	return queens

	// 回溯
	//queens := make([][]string, 0)
	//pie := make([]bool, n<<1-1)	// 可以将 数组 换为 map
	//na := make([]bool, n<<1-1)
	//column := make([]bool, n)
	//solve := make([]int, n)
	//dfsQueens(&queens, 0, solve, pie, na, column, n)
	//return queens
}

func dfsBits(pie, na, c int, row int, solve []int, queens *[][]string, n int) {
	if row == n {
		getQueen(solve, queens, n)
		return
	}
	nBits := 1<<n - 1	// 因为后面要进行移位运算，所以只能采取：传参时用 1 记录已被占位，0 记录可以使用。而使用时，反过来方便取最后一位1，等运算
	ok := ^(pie | na | c) & nBits
	for ok > 0 {
		last := ok & (-ok)
		//col := math.Log2(float64(last))	// 使用API
		col := 0
		for testBits := 1 << (n - 1); testBits&last == 0; testBits >>= 1 {
			col++ // 顺序是倒过来遍历的
		}
		solve[row] = col
		dfsBits((pie|last)<<1, (na|last)>>1, c|last, row+1, solve, queens, n)
		ok ^= last // 因为并没有改变 pie na c 的值，所以不需要回溯
	}
}

func dfsQueens(queens *[][]string, row int, solve []int, pie []bool, na []bool, column []bool, n int) {
	if row == n {
		getQueen(solve, queens, n)
		return
	}
	for col := 0; col < n; col++ {
		if pie[col+row] || na[col+n-1-row] || column[col] {
			continue
		}
		solve[row] = col
		pie[col+row], na[col+n-1-row], column[col] = true, true, true
		dfsQueens(queens, row+1, solve, pie, na, column, n)
		pie[col+row], na[col+n-1-row], column[col] = false, false, false
	}
}

func getQueen(solve []int, queens *[][]string, n int) {
	queen := make([]string, n)
	for i, col := range solve {
		strRow := make([]byte, n)
		for j := 0; j < n; j++ {
			if j == col {
				strRow[j] = 'Q'
			} else {
				strRow[j] = '.'
			}
		}
		queen[i] = string(strRow)
	}
	*queens = append(*queens, queen)
}

//leetcode submit region end(Prohibit modification and deletion)
