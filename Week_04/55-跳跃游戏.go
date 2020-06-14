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
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{3, 2, 1, 0, 0, 100, 4}
	nums := []int{1, 2, 3}
	jump := canJump(nums)
	fmt.Println(jump)
}

//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	// 贪心：从后
	target := len(nums) - 1
	for i := target - 1; i >= 0; i-- {
		if nums[i]+i >= target {
			target = i
		}
	}
	return target == 0

	// 贪心：从前
	//can := 0
	//for i := 0; i < len(nums)-1; i++ {
	//	if can < i {
	//		return false
	//	}
	//	if nums[i]+i > can {
	//		can = nums[i] + i
	//		if can >= len(nums)-1 {
	//			return true
	//		}
	//	}
	//}
	//return false

	// dp
	//need := 0
	//for i := len(nums) - 2; i >= 0; i-- {
	//	if nums[i+1] >= need { // 上一个数，是否大于它需要的步数，大于等于：当前步至少需要 1，小于：++
	//		need = 1
	//	} else {
	//		need++
	//	}
	//}
	//return nums[0] >= need
}

//leetcode submit region end(Prohibit modification and deletion)
