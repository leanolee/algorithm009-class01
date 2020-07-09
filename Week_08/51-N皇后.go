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
	n := 4
	queens := solveNQueens(n)
	fmt.Println(queens)
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	// 位运算
	ans, pie, na, c := make([][]string, 0), 0, 0, 0
	nQueensDFS(n, &ans, make([]int, n), 0, pie, na, c)
	return ans

	// 数组：dfs + 回溯
	//pie, na, col := make([]bool, n<<1-1), make([]bool, n<<1-1), make([]bool, n)
	//ans := make([][]string, 0)
	//solveNQueensDFS(n, &ans, make([]int, n), 0, pie, na, col)
	//return ans
}

func nQueensDFS(n int, ans *[][]string, solve []int, i int, pie int, na int, c int) {
	if i == n {
		*ans = append(*ans, getQueen(n, solve))
		return
	}
	bits := ^(pie | na | c) & (1<<n - 1)
	for bits > 0 {
		last := bits & -bits
		idx := 0
		for test := 1; test&last == 0; test <<= 1 {
			idx++
		}
		solve[i] = idx
		nQueensDFS(n, ans, solve, i+1, (pie|last)<<1, (na|last)>>1, c|last)
		bits ^= last
	}
}

func solveNQueensDFS(n int, ans *[][]string, solve []int, r int, pie []bool, na []bool, col []bool) {
	if r == n {
		*ans = append(*ans, getQueen(n, solve))
		return
	}
	for c := 0; c < n; c++ {
		if pie[r+c] || na[n-r+c-1] || col[c] {
			continue
		}
		solve[r] = c
		pie[r+c], na[n-r+c-1], col[c] = true, true, true
		solveNQueensDFS(n, ans, solve, r+1, pie, na, col)
		pie[r+c], na[n-r+c-1], col[c] = false, false, false
	}
}

func getQueen(n int, solve []int) []string {
	queen := make([]string, n)
	for i, v := range solve {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if j == v {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		queen[i] = string(row)
	}
	return queen
}

//leetcode submit region end(Prohibit modification and deletion)
