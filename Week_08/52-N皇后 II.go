//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。 
//
// 示例: 
//
// 输入: 4
//输出: 2
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
// 
//
// 
//
// 提示： 
//
// 
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或七步
//，可进可退。（引用自 百度百科 - 皇后 ） 
// 
// Related Topics 回溯算法
package main

import "fmt"

func main() {
	n := 4
	queens := totalNQueens(n)
	fmt.Println(queens)
}

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	// 位运算：
	return totalNQueensDFS(n, 0, 0, 0, 0)

	// dfs：
	//pie, na, c := make([]bool, n<<1-1), make([]bool, n<<1-1), make([]bool, n)
	//return totalNQueensDfs(n, 0, pie, na, c)
}

func totalNQueensDFS(n int, i int, pie, na, c int) int {
	if i == n {
		return 1
	}
	bits, ans := ^(pie|na|c)&(1<<n-1), 0
	for bits > 0 {
		last := bits & -bits
		ans += totalNQueensDFS(n, i+1, (pie|last)<<1, (na|last)>>1, c|last)
		bits ^= last
	}
	return ans
}

func totalNQueensDfs(n int, i int, pie []bool, na []bool, c []bool) int {
	if i == n {
		return 1
	}
	ans := 0
	for j := 0; j < n; j++ {
		if pie[i+j] || na[n-i-1+j] || c[j] {
			continue
		}
		pie[i+j], na[n-i-1+j], c[j] = true, true, true
		ans += totalNQueensDfs(n, i+1, pie, na, c)
		pie[i+j], na[n-i-1+j], c[j] = false, false, false
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
