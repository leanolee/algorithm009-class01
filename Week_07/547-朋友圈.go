//班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 
//的朋友。所谓的朋友圈，是指所有朋友的集合。 
//
// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你
//必须输出所有学生中的已知的朋友圈总数。 
//
// 示例 1: 
//
// 
//输入: 
//[[1,1,0],
// [1,1,0],
// [0,0,1]]
//输出: 2 
//说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
//第2个学生自己在一个朋友圈。所以返回2。
// 
//
// 示例 2: 
//
// 
//输入: 
//[[1,1,0],
// [1,1,1],
// [0,1,1]]
//输出: 1
//说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
// 
//
// 注意： 
//
// 
// N 在[1,200]的范围内。 
// 对于所有学生，有M[i][i] = 1。 
// 如果有M[i][j] = 1，则有M[j][i] = 1。 
// 
// Related Topics 深度优先搜索 并查集
package main

import "fmt"

func main() {
	M := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}} // 3
	//M := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} // 3
	//M := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}} // 1
	num := findCircleNum(M)
	fmt.Println(num)
}

/*
思路：
	参考：leetcode-200：岛屿数量
并查集：
	p、q 都是指索引
	find(parent []int, p int) int：查找代表
	union(parent []int, p, q int)：合并
*/
//leetcode submit region begin(Prohibit modification and deletion)
func findCircleNum(M [][]int) int {
	// 并查集
	find := func(parent []int, p int) int {
		for p != parent[p] {
			parent[p], p = parent[parent[p]], parent[parent[p]] // 有路径压缩的效果：合并了 p 和 parent[p]
		}
		return p
	}
	union := func(parent []int, p, q int) {
		rootP, rootQ := find(parent, p), find(parent, q)
		if rootP != rootQ {
			parent[rootP] = rootQ //合并
		}
	}
	count, n := 0, len(M)
	disjoint := make([]int, n)
	for i, _ := range disjoint {
		disjoint[i] = i // 初始化：自己是自己的朋友
	}
	for i := 1; i < n; i++ { // 可以从 1 开始
		for j := 0; j < i; j++ { // 缩小遍历
			if M[i][j] == 1 { // 有朋友关系，则合并
				union(disjoint, i, j)
			}
		}
	}
	for i, d := range disjoint {
		if i == d {
			count++
		}
	}
	return count

	// DFS：连锁反应
	//count, r := 0, len(M)
	//visited := make([]bool, r) // 记录有没有被记录过朋友关系
	//var dfs func(i int)
	//dfs = func(i int) {	// 特殊的 dfs
	//	for j := 0; j < r; j++ { // 遍历 i 和其他人的关系
	//		if M[i][j] == 1 && !visited[j] {
	//			visited[j] = true
	//			dfs(j) // 遍历下一个人
	//		}
	//	}
	//}
	//for i := 0; i < r; i++ { // 遍历每个人
	//	if !visited[i] {
	//		dfs(i)
	//		count++
	//	}
	//}
	//return count

	// BFS
}

//leetcode submit region end(Prohibit modification and deletion)
