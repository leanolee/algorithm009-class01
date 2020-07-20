//你的赛车起始停留在位置 0，速度为 +1，正行驶在一个无限长的数轴上。（车也可以向负数方向行驶。） 
//
// 你的车会根据一系列由 A（加速）和 R（倒车）组成的指令进行自动驾驶 。 
//
// 当车得到指令 "A" 时, 将会做出以下操作： position += speed, speed *= 2。 
//
// 当车得到指令 "R" 时, 将会做出以下操作：如果当前速度是正数，则将车速调整为 speed = -1 ；否则将车速调整为 speed = 1。 (当前所
//处位置不变。) 
//
// 例如，当得到一系列指令 "AAR" 后, 你的车将会走过位置 0->1->3->3，并且速度变化为 1->2->4->-1。 
//
// 现在给定一个目标位置，请给出能够到达目标位置的最短指令列表的长度。 
//
// 示例 1:
//输入: 
//target = 3
//输出: 2
//解释: 
//最短指令列表为 "AA"
//位置变化为 0->1->3
// 
//
// 示例 2:
//输入: 
//target = 6
//输出: 5
//解释: 
//最短指令列表为 "AAARA"
//位置变化为 0->1->3->7->7->6
// 
//
// 说明: 
//
// 
// 1 <= target（目标位置） <= 10000。 
// 
// Related Topics 堆 动态规划
package main

import (
	"fmt"
	"math"
)

func main() {
	//target := 77	// 13
	target := 54 // 15
	i := racecar(target)
	fmt.Println(i)
	//	0 1 4 2 5 7 5 3 6 8 7 10 7 9 6 4 7 9  8 11 12 10 9 12 9 11 13 11 8 10 7 5 8 10 9 12 13 11 10 13 15 14 15 13 14 12 11 14 11 13 16 14 17 14 15 	// 54
	//	13 10 12 14 12 9 11 8 6 9 11 10 13 14 12 11 14 16 15 16	// 74
	//  0 1 4 2 5 8 5 3 6 9 7 10 7 9 6 4 7 10 8 11 14 11 9 12 9 11 14 11 8 10 7 5 8 11 9 12 15 12 10 13 16 14 17 14 16 13 11 14 11 13 16 14 17 14 16 	// 54
	//  13 10 12 15 12 9 11 8 6 9 12 10 13 16 13 11 14 17 15 18	// 74
}

/*
dp：详情请见草稿
*/
//leetcode submit region begin(Prohibit modification and deletion)
func racecar(target int) int {
	// lc：最短路

	// lc：dp
	dp := make([]int, target+1)
	dp[1] = 1
	for i, s, p := 2, 2, 1; i <= target; i++ {
		dp[i] = math.MaxInt32
		if p+s > i {
			for j := 0; j < dp[p]; j++ {
				dp[i] = min(dp[i], dp[i-p+1<<j-1]+dp[p]+j+2)
			}
			dp[i] = min(dp[i], dp[p+s-i]+2+dp[p])
			//dp[i] = min(dp[i-p], dp[p+s-i]) + 2 + dp[p]
		} else {
			dp[i] = dp[p] + 1
			p += s
			s <<= 1
		}
	}
	fmt.Println(dp)
	return dp[target]

	// 个人方法：dp + 位运算，未通过，因为没考虑折返的情况
	//dp := make([]int, target+1)
	//dp[1] = 1
	//for i, s, p := 2, 2, 1; i <= target; i++ {
	//	if p+s > i {
	//		dp[i] = min(dp[i-p], dp[p+s-i]) + 2 + dp[p]	// 考虑折返
	//	} else {
	//		dp[i] = dp[p] + 1
	//		p += s
	//		s <<= 1
	//	}
	//}
	//fmt.Println(dp)
	//return dp[target]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
一种快捷求位数的方法，详情见 algorithms...
*/
func getBitCount(i int) int {
	n := 0
	if i>>16 != 0 {
		n += 16
		i >>= 16
	}
	if i>>8 != 0 {
		n += 8
		i >>= 8
	}
	if i>>4 != 0 {
		n += 4
		i >>= 4
	}
	if i>>2 != 0 {
		n += 2
		i >>= 2
	}
	if i>>1 != 0 {
		n += 1
		i >>= 1
	}
	return n + i
}

//leetcode submit region end(Prohibit modification and deletion)
