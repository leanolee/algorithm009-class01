//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。 
//
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。 
//
// 示例: 
//
// X X X X
//X O O X
//X X O X
//X O X X
// 
//
// 运行你的函数后，矩阵变为： 
//
// X X X X
//X X X X
//X X X X
//X O X X
// 
//
// 解释: 
//
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被
//填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。 
// Related Topics 深度优先搜索 广度优先搜索 并查集
package main

import (
	"container/list"
	"fmt"
)

func main() {
	//board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	//board := [][]byte{
	//	{'O', 'X', 'X', 'O', 'X'},
	//	{'X', 'O', 'O', 'X', 'O'},
	//	{'X', 'O', 'X', 'O', 'X'},
	//	{'O', 'X', 'O', 'O', 'O'},
	//	{'X', 'X', 'O', 'X', 'O'}} // 重点：当这种情况时，i=3，board[2][3]被包围了，回滚不了
	board := [][]byte{
		{'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'O', 'O'}} // 重点：当这种情况时，i=3，board[2][3]被包围了，回滚不了
	solve(board)
	for _, row := range board {
		for _, c := range row {
			fmt.Printf("%c  ", c)
		}
		fmt.Println()
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte) {
	// 并查集
	//if len(board) == 0 {
	//	return
	//}
	//m, n := len(board), len(board[0])
	//parent, di, dj := make([]int, m*n+1), [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	//for i, _ := range parent {
	//	parent[i] = i
	//}
	//find := func(parent []int, p int) int {
	//	for p != parent[p] {
	//		parent[p], p = parent[parent[p]], parent[parent[p]]
	//	}
	//	return p
	//}
	//union := func(parent []int, p, q int) {
	//	rootP, rootQ := find(parent, p), find(parent, q)
	//	if rootP != rootQ {
	//		parent[rootP] = rootQ
	//	}
	//}
	//var dfs func(i, j int)
	//dfs = func(i, j int) {
	//	fmt.Println(i, j)
	//	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 'X' || find(parent, i*n+j) == len(parent)-1 {
	//		return
	//	}
	//	union(parent, i*n+j, len(parent)-1)
	//	for k := 0; k < 4; k++ {
	//		dfs(i+di[k], j+dj[k])
	//	}
	//}
	//for j := 0; j < n; j++ {
	//	if board[0][j] == 'O' {
	//		dfs(0, j)
	//	}
	//}
	//for i := 1; i < m; i++ {
	//	if board[i][0] == 'O' {
	//		dfs(i, 0)
	//	}
	//}
	//for i := 1; i < m; i++ {
	//	if board[i][n-1] == 'O' {
	//		dfs(i, n-1)
	//	}
	//}
	//for j := 1; j < n-1; j++ {
	//	if board[m-1][j] == 'O' {
	//		dfs(m-1, j)
	//	}
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if board[i][j] == 'O' && find(parent, i*n+j) != len(parent)-1 {
	//			board[i][j] = 'X'
	//		}
	//	}
	//}

	// BFS：不会OOM，重点：四连通 添加了很多重复的元素到 队列 中
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	queue, di, dj := list.New(), [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	visited := make([][]bool, m) // BFS oom 的关键不是队列中的元素多了，而是 四连通 添加了很多重复的元素到 队列 中
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	bfs := func(i, j int) {
		queue.PushBack([2]int{i, j}) // 采用 list 作为队列
		visited[i][j] = true
		for queue.Len() > 0 {
			curr := queue.Remove(queue.Front()).([2]int) // 新写法
			for k := 0; k < 4; k++ {
				nI, nJ := curr[0]+di[k], curr[1]+dj[k]
				if nI < 0 || nI >= m || nJ < 0 || nJ >= n || board[nI][nJ] == 'X' || visited[nI][nJ] {
					continue
				}
				queue.PushBack([2]int{nI, nJ})
				visited[nI][nJ] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				bfs(i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}

	// BFS：难点，out of memory（太多O），重点：见上一种 BFS
	//if len(board) == 0 {
	//	return
	//}
	//m, n := len(board), len(board[0])
	//queue, di, dj := make([]Point, 0), [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	//var bfs func()
	//bfs = func() {
	//	//for idx := 0; idx < len(queue); idx++ {	// out of memory（太多O），
	//	for len(queue) > 0 { // out of memory（太多O），那么就要换种写法
	//		length := len(queue)
	//		for idx := 0; idx < length; idx++ {
	//			i, j := queue[idx].i, queue[idx].j
	//			board[i][j] = '!'
	//			for k := 0; k < 4; k++ {
	//				nI, nJ := i+di[k], j+dj[k]
	//				if nI < 0 || nI >= m || nJ < 0 || nJ >= n || board[nI][nJ] == 'X' || board[nI][nJ] == '!' {
	//					continue
	//				}
	//				queue = append(queue, Point{nI, nJ})	// 作为 四连通 添加重复元素的案例
	//			}
	//		}
	//		queue = queue[length:]
	//	}
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if board[i][j] == 'O' && (i == 0 || i == m-1 || j == 0 || j == n-1) {
	//			queue = append(queue, Point{i, j})
	//			bfs()
	//		}
	//	}
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		switch board[i][j] {
	//		case '!':
	//			board[i][j] = 'O'
	//		case 'O':
	//			board[i][j] = 'X'
	//		}
	//	}
	//}

	// DFS
	//m := len(board) - 1
	//if m == -1 {
	//	return
	//}
	//n := len(board[0]) - 1
	//dx, dy := [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	//var dfs func(i, j int)
	//dfs = func(i, j int) {
	//	if i < 0 || i > m || j < 0 || j > n || board[i][j] == 'X' || board[i][j] == '!' {
	//		return
	//	}
	//	board[i][j] = '!'
	//	for k := 0; k < 4; k++ {
	//		dfs(i+dx[k], j+dy[k])
	//	}
	//}
	//for i := 0; i <= m; i++ {
	//	for j := 0; j <= n; j++ {
	//		if i == 0 || i == m || j == 0 || j == n && board[i][j] == 'O' { // 处理边界
	//			dfs(i, j)
	//		}
	//	}
	//}
	//for i := 0; i <= m; i++ {
	//	for j := 0; j <= n; j++ {
	//		switch board[i][j] { // 处理中央和边界
	//		case '!':
	//			board[i][j] = 'O'
	//		case 'O':
	//			board[i][j] = 'X'
	//		}
	//	}
	//}

	// DFS：失败
	//if len(board) == 0 {
	//	return
	//}
	//r, c := len(board), len(board[0])
	//dx, dy := [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	//visited := make([]bool, r*c) // 加 visited 也解决不了“封闭”的bug
	//var dfs func(i, j int)       // 有返回值的情况，行不通，因为当 [i,j] 3面都是 X 时，进来的地方，进来的地方和 [i,j] 互相依赖，无法判断
	//dfs = func(i, j int) {
	//	fmt.Println(i, j)
	//	if board[i][j] == 'X' { // 这里也要判断 visited[i*c+j]，因为前面可能提前触到边界返回了
	//		return
	//	}
	//	visited[i*c+j] = true
	//	//ans := false
	//	for k := 0; k < 4; k++ {
	//		nI, nJ := i+dx[k], j+dy[k]
	//		if nI < 0 || nI == r || nJ < 0 || nJ == c {
	//			return
	//		}
	//		if visited[i*c+j] {
	//			continue
	//		}
	//		dfs(nI, nJ)
	//	}
	//	board[i][j] = 'X'
	//}
	//for i := 0; i < r; i++ {
	//	for j := 0; j < c; j++ {
	//		if board[i][j] == 'O' && !visited[i*c+j] {
	//			dfs(i, j)
	//		}
	//	}
	//}
}

type Point struct {
	i int
	j int
}

//leetcode submit region end(Prohibit modification and deletion)
