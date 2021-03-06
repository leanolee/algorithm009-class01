//亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。 
//
// 游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。 
//
// 亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。
// 
//
// 假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。 
//
// 
//
// 示例： 
//
// 输入：[5,3,4,5]
//输出：true
//解释：
//亚历克斯先开始，只能拿前 5 颗或后 5 颗石子 。
//假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
//如果李拿走前 3 颗，那么剩下的是 [4,5]，亚历克斯拿走后 5 颗赢得 10 分。
//如果李拿走后 5 颗，那么剩下的是 [3,4]，亚历克斯拿走后 4 颗赢得 9 分。
//这表明，取前 5 颗石子对亚历克斯来说是一个胜利的举动，所以我们返回 true 。
// 
//
// 
//
// 提示： 
//
// 
// 2 <= piles.length <= 500 
// piles.length 是偶数。 
// 1 <= piles[i] <= 500 
// sum(piles) 是奇数。 
// 
// Related Topics 极小化极大 数学 动态规划
package main

import "fmt"

func main() {
	piles := []int{3, 9, 1, 2}
	game := stoneGame(piles)
	fmt.Println(game)
}

/*
博弈DP：https://leetcode-cn.com/problems/stone-game/solution/jie-jue-bo-yi-wen-ti-de-dong-tai-gui-hua-tong-yong/
*/
//leetcode submit region begin(Prohibit modification and deletion)
func stoneGame(piles []int) bool {
	// 数学：关键，piles.length 是偶数
	//return true

	// dp：
	n := len(piles)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
	}
	for i := n - 1; i >= 0; i-- { // 倒着遍历半个矩阵
		dp[i][i][0] = piles[i]       // 默认不管是奇数还是偶数个，第一个人先选
		for j := i + 1; j < n; j++ { // i+1
			/*
				dp[i][j].fir = max(piles[i] + dp[i+1][j].sec, piles[j] + dp[i][j-1].sec)
				dp[i][j].fir = max(    选择最左边的石头堆     ,     选择最右边的石头堆     )
				if 先手选择左边:
				    dp[i][j].sec = dp[i+1][j].fir
				if 先手选择右边:
				    dp[i][j].sec = dp[i][j-1].fir
			*/
			l, r := piles[i]+dp[i+1][j][1], piles[j]+dp[i][j-1][1]
			if l < r {
				dp[i][j][0], dp[i][j][1] = r, dp[i][j-1][0]
			} else {
				dp[i][j][0], dp[i][j][1] = l, dp[i+1][j][0]
			}
		}
	}
	return dp[0][n-1][0] > dp[0][n-1][1]

	//dp := make([][]int, n+2)
	//for i := 0; i <= n+1; i++ {
	//	dp[i] = make([]int, n+2)
	//}
	//for size := 1; size <= n; size++ { // size
	//	for i := 0; i+size <= n; i++ {
	//		j := i + size - 1
	//		if (j-i+1)&1 == 1 { // 奇数：轮到第二人取，直接减去第 1 人的数
	//			dp[i+1][j+1] = min(dp[i+2][j]-piles[i], dp[i+1][j]-piles[j])
	//		} else { // 偶数
	//			dp[i+1][j+1] = max(dp[i+2][j]+piles[i], dp[i+1][j]+piles[j])
	//		}
	//	}
	//}
	//return dp[1][n] > 0
}

//leetcode submit region end(Prohibit modification and deletion)
