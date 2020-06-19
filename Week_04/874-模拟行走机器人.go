//机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令： 
//
// 
// -2：向左转 90 度 
// -1：向右转 90 度 
// 1 <= x <= 9：向前移动 x 个单位长度 
// 
//
// 在网格上有一些格子被视为障碍物。 
//
// 第 i 个障碍物位于网格点 (obstacles[i][0], obstacles[i][1]) 
//
// 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。 
//
// 返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。 
//
// 
//
// 示例 1： 
//
// 输入: commands = [4,-1,3], obstacles = []
//输出: 25
//解释: 机器人将会到达 (3, 4)
// 
//
// 示例 2： 
//
// 输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
//输出: 65
//解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处
// 
//
// 
//
// 提示： 
//
// 
// 0 <= commands.length <= 10000 
// 0 <= obstacles.length <= 10000 
// -30000 <= obstacle[i][0] <= 30000 
// -30000 <= obstacle[i][1] <= 30000 
// 答案保证小于 2 ^ 31 
// 
// Related Topics 贪心算法
package main

import (
	"fmt"
)

func main() {
	commands := []int{7, -2, -2, 7, 5}
	obstacles := [][]int{{-3, 2}, {-2, 1}, {0, 1}, {-2, 4}, {-1, 0}, {-2, -3}, {0, -3}, {4, 4}, {-3, 3}, {2, 2}}
	//commands := []int{-2, -1, -2, 3, 7}
	//obstacles := [][]int{{1, -3}, {2, -3}, {4, 0}, {-2, 5}, {-5, 2}, {0, 0}, {4, -4}, {-2, -5}, {-1, -2}, {0, 2}}
	sim := robotSim(commands, obstacles)
	fmt.Println(sim)

	//i := 0
	//for num := 1; num <= 30000; {
	//	i++
	//	num <<= 1
	//}
	//fmt.Println(i) // 15
}

//leetcode submit region begin(Prohibit modification and deletion)
func robotSim(commands []int, obstacles [][]int) int {
	// leetcode
	getPoint := func(x, y int) int {
		return (x+30000)<<16 + 30000 + y
	}
	memo := make(map[int]bool)
	for _, obs := range obstacles {
		//memoX[obs[0]] = obs[1]	// 被覆盖了，所以失败了
		memo[getPoint(obs[0], obs[1])] = true
	}
	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	x, y, direct, max := 0, 0, 0, 0
	for _, c := range commands {
		switch c {
		case -2:
			direct--
			if direct == -1 {
				direct = 3
			}
		case -1:
			direct++
			if direct == 4 {
				direct = 0
			}
		default:
			for i := 0; i < c; i++ {
				pathX, pathY := x+dx[direct], y+dy[direct]
				if memo[getPoint(pathX, pathY)] {
					break
				}
				x, y = pathX, pathY
			}
			curr := x*x + y*y
			if curr > max {
				max = curr
			}
		}
	}
	return max

	// leetcode-cn
	//dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	//memo := make(map[int]bool)
	//getPoint := func(x, y int) int {
	//	return (x+30000)<<16 + y + 30000
	//}
	//for _, obs := range obstacles {
	//	memo[getPoint(obs[0], obs[1])] = true
	//}
	//x, y, direct := 0, 0, 0
	//max := 0
	//for _, c := range commands {
	//	switch c {
	//	case -2:
	//		direct = (direct + 3) % 4
	//	case -1:
	//		direct = (direct + 1) % 4
	//	default:
	//		for i := 1; i <= c; i++ {
	//			pathX, pathY := x+dx[direct], y+dy[direct]
	//			if memo[getPoint(pathX, pathY)] {
	//				break
	//			}
	//			x, y = pathX, pathY
	//		}
	//		curr := x*x + y*y
	//		if curr > max {
	//			max = curr
	//		}
	//	}
	//}
	//return max

	// 个人写法
	//x, y, direction := 0, 0, 0 // direction：0 1 2 3 -> 北 东 南 西
	//hasObstacle := func(fix, to int, fixIdx, toIdx int, tar int) int { // xOrY：0 1 -> x y
	//	for _, obs := range obstacles {
	//		if obs[fixIdx] == fix {
	//			obsVal := obs[toIdx]
	//			if to < obsVal && obsVal <= tar {
	//				return obsVal - 1
	//			} else if to > obsVal && obsVal >= tar {
	//				return obsVal + 1
	//			}
	//		}
	//	}
	//	return tar
	//}
	//max := 0
	//for _, m := range commands {
	//	switch m {
	//	case -2:
	//		direction = (direction + 3) % 4
	//	case -1:
	//		direction = (direction + 1) % 4
	//	default: // 1 <= m <= 9
	//		switch direction {
	//		case 0:
	//			y = hasObstacle(x, y, 0, 1, y+m)
	//		case 1:
	//			x = hasObstacle(y, x, 1, 0, x+m)
	//		case 2:
	//			y = hasObstacle(x, y, 0, 1, y-m)
	//		case 3:
	//			x = hasObstacle(y, x, 1, 0, x-m)
	//		}
	//	}
	//	curr := x*x + y*y
	//	if curr > max {
	//		max = curr
	//	}
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
