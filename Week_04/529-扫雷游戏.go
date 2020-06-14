//让我们一起来玩扫雷游戏！ 
//
// 给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）
//地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。 
//
// 现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板： 
//
// 
// 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。 
// 如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的方块都应该被递归地揭露。 
// 如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。 
// 如果在此次点击中，若无更多方块可被揭露，则返回面板。 
// 
//
// 
//
// 示例 1： 
//
// 输入: 
//
//[['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'M', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E']]
//
//Click : [3,0]
//
//输出: 
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
//解释:
//
// 
//
// 示例 2： 
//
// 输入: 
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
//Click : [1,2]
//
//输出: 
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'X', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
//解释:
//
// 
//
// 
//
// 注意： 
//
// 
// 输入矩阵的宽和高的范围为 [1,50]。 
// 点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。 
// 输入面板不会是游戏结束的状态（即有地雷已被挖出）。 
// 简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。 
// Related Topics 深度优先搜索 广度优先搜索
package main

import "fmt"

func main() {
	board := [][]byte{{'B', '1', 'E', '1', 'B'}, {'B', '1', 'M', '1', 'B'}, {'B', '1', '1', '1', 'B'}, {'B', 'B', 'B', 'B', 'B'}}
	click := []int{3, 0}
	res := updateBoard(board, click)
	strs := make([][]string, len(res))
	for i := 0; i < len(res); i++ {
		strs[i] = make([]string, len(res[0]))
		for j := 0; j < len(res[0]); j++ {
			strs[i][j] = string(res[i][j])
		}
	}
	fmt.Println(strs)
}

/*
1.M：地雷
2.E：B
	dfs
如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的方块都应该被递归地揭露。
如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
如果在此次点击中，若无更多方块可被揭露，则返回面板。
*/
//leetcode submit region begin(Prohibit modification and deletion)
func updateBoard(board [][]byte, click []int) [][]byte {
	// DFS
	row, col := len(board), len(board[0])
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	var dfs func(int, int)
	dfs = func(i, j int) {
		var mCount byte = '0'
		for r := -1; r < 2; r++ {
			for c := -1; c < 2; c++ {
				if r|c == 0 {
					continue
				}
				x, y := i+r, j+c
				if row > x && x >= 0 && col > y && y >= 0 && board[x][y] == 'M' {
					mCount++
				}
			}
		}
		if mCount > '0' {
			board[i][j] = mCount
		} else {
			board[i][j] = 'B'
			for r := -1; r < 2; r++ {
				for c := -1; c < 2; c++ {
					if r|c == 0 {
						continue
					}
					x, y := i+r, j+c
					if row > x && x >= 0 && col > y && y >= 0 && board[x][y] == 'E' {
						dfs(x, y)
					}
				}
			}
		}
	}
	dfs(x, y)
	return board

	// DFS：个人写法
	//row := len(board)
	//col := len(board[0])
	//x, y := click[0], click[1]
	//if board[x][y] == 'M' {
	//	board[x][y] = 'X'
	//	return board
	//}
	//dfs(board, row, col, x, y)
	//return board
}

func dfs(board [][]byte, row int, col int, i int, j int) byte {
	if i < 0 || j < 0 || i >= row || j >= col || board[i][j] == 'B' {
		return 0
	}
	if board[i][j] == 'M' { // 先看是否是炸弹
		return 1
	}
	if board[i][j] != 'E' {
		return board[i][j] - '0'
	}
	var mCount byte = 0
	mCount += dfs(board, row, col, i-1, j-1)
	mCount += dfs(board, row, col, i-1, j)
	mCount += dfs(board, row, col, i-1, j+1)
	mCount += dfs(board, row, col, i, j-1)
	mCount += dfs(board, row, col, i, j+1)
	mCount += dfs(board, row, col, i+1, j-1)
	mCount += dfs(board, row, col, i+1, j)
	mCount += dfs(board, row, col, i+1, j+1)
	if mCount > 0 {
		board[i][j] = mCount + '0'
	} else {
		board[i][j] = 'B'
	}
	return mCount
}

//leetcode submit region end(Prohibit modification and deletion)
