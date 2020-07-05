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
	n := 10
	queens := totalNQueens(n)
	fmt.Println(queens)
}

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	// DFS + 位运算 + 剪枝
	var dfs func(i int, pie, na, c int) int
	dfs = func(i int, pie, na, c int) int {
		if i == n {
			return 1
		}
		memo, ans := ^(pie|na|c)&(1<<n-1), 0 // 1：可用
		for memo > 0 {
			last := memo & -memo // 取最后一个 1
			memo ^= last
			ans += dfs(i+1, (pie|last)<<1, (na|last)>>1, c|last)
		}
		return ans
	}
	return dfs(0, 0, 0, 0)

	// DFS + 数组 + 回溯 + 剪枝
	//pie, na, c := make([]bool, n<<1-1), make([]bool, n<<1-1), make([]bool, n) // x+y n-x+y
	//var dfs func(i int, pie []bool, na []bool, c []bool) int
	//dfs = func(i int, pie []bool, na []bool, c []bool) int {
	//	if i == n {
	//		return 1
	//	}
	//	ans := 0
	//	for j := 0; j < n; j++ {
	//		if pie[i+j] || na[n-i+j-1] || c[j] {
	//			continue
	//		}
	//		pie[i+j], na[n-i+j-1], c[j] = true, true, true
	//		ans += dfs(i+1, pie, na, c)
	//		pie[i+j], na[n-i+j-1], c[j] = false, false, false
	//	}
	//	return ans
	//}
	//return dfs(0, pie, na, c)
}

//leetcode submit region end(Prohibit modification and deletion)
