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
	queens := make([][]string, 0)
	dfs(n, &queens, nil, 0, 0, 0, 0)
	return queens
}

func dfs(n int, ans *[][]string, solve []int, row int, pie, na, c int) {
	if row == n {
		*ans = append(*ans, getQueen(n, solve))
		return
	}
	ok := ^(pie | na | c) & (1<<n - 1)
	for ok > 0 {
		last := ok & -ok
		idx := 0
		for test := 1; test&last == 0; test <<= 1 {
			idx++
		}
		dfs(n, ans, append(solve, idx), row+1, (pie|last)<<1, (na|last)>>1, c|last)
		ok ^= last
	}
}

func getQueen(n int, solve []int) []string {
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if solve[i] == j {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		ans[i] = string(row)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
