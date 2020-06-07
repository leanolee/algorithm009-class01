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

import "fmt"

func main() {
	n := 9
	queens := solveNQueens(n)
	fmt.Println(queens, len(queens))
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	// 位运算
	queens := make([][]string, 0)
	dfsBits(n, &queens, make([]int, n), 0, 0, 0, 0)
	return queens

	// DFS：回溯
	//queens := make([][]string, 0)
	//pie := make([]bool, n<<1+1)
	//na := make([]bool, n<<1+1)
	//c := make([]bool, n)
	//dfsQueens(n, &queens, make([]int, n), 0, pie, na, c)
	//return queens
}

func dfsBits(n int, queens *[][]string, solve []int, row int, pie, na, c int) {
	if row == n {
		*queens = append(*queens, getQueen(n, solve))
		return
	}
	ok := ^(pie | na | c) & (1<<n - 1)
	for ok > 0 {
		last := ok & -ok // 取出最后一位
		idx := 0
		for idxTest := 1 << (n - 1); idxTest&last == 0; idxTest >>= 1 {
			idx++
		}
		solve[row] = idx
		dfsBits(n, queens, solve, row+1, (pie^last)<<1, (na^last)>>1, c^last)
		ok ^= last
	}
}

func dfsQueens(n int, queens *[][]string, solve []int, row int, pie, na, c []bool) {
	if row == n {
		*queens = append(*queens, getQueen(n, solve))
		return
	}
	for col := 0; col < n; col++ {
		if pie[row+col] || na[n+col-1-row] || c[col] {
			continue
		}
		solve[row] = col
		pie[row+col], na[n+col-1-row], c[col] = true, true, true
		dfsQueens(n, queens, solve, row+1, pie, na, c)
		pie[row+col], na[n+col-1-row], c[col] = false, false, false
	}
}

func getQueen(n int, ints []int) []string {
	queen := make([]string, n)
	for i, q := range ints {
		s := make([]byte, n)
		for j := 0; j < n; j++ {
			if j != q {
				s[j] = '.'
			} else {
				s[j] = 'Q'
			}
		}
		queen[i] = string(s)
	}
	return queen
}

//leetcode submit region end(Prohibit modification and deletion)
