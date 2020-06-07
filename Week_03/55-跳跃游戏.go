//给定一个非负整数数组，你最初位于数组的第一个位置。 
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。 
//
// 判断你是否能够到达最后一个位置。 
//
// 示例 1: 
//
// 输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
// 
//
// 示例 2: 
//
// 输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
// 
// Related Topics 贪心算法 数组
package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{0}
	jump := canJump(nums)
	fmt.Println(jump)
}

//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	// 贪心：优化
	//can := len(nums) - 1
	//for i := can - 1; i >= 0; i-- {
	//	if nums[i]+i >= can {	// 当前值 + 当前位置 >= 上次可以的位置
	//		can = i
	//	}
	//}
	//return can == 0

	// 贪心：
	mostFar := 0 // 能跳到最远的索引
	for i := 0; i < len(nums); i++ {
		if i > mostFar { // 已经落后了
			return false
		}
		if nums[i]+i >= mostFar { // 新记录
			mostFar = nums[i] + i
			if mostFar >= len(nums)-1 {
				return true
			}
		}
	}
	return false

	// dp：优化
	//n := len(nums)
	//if n == 1 {
	//	return true
	//}
	//needNum := 1
	//for i := n - 2; i >= 0; i-- {
	//	if nums[i] >= needNum {
	//		needNum = 1
	//	} else {
	//		needNum++
	//	}
	//}
	//return nums[0] >= needNum

	// dp：
	//n := len(nums)
	//if n == 1 {
	//	return true
	//}
	//dp := make([]int, n)
	//dp[n-1] = 1
	//for i := n - 2; i >= 0; i-- {
	//	if nums[i] >= dp[i+1] {
	//		dp[i] = 1
	//	} else {
	//		dp[i] = dp[i+1] + 1
	//	}
	//}
	//return nums[0] >= dp[0]

	// 暴力法
	//return false
}

//leetcode submit region end(Prohibit modification and deletion)
