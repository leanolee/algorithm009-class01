//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。 
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//] 
// Related Topics 位运算 数组 回溯算法
package main

import "fmt"

func main() {
	//[[],[9],[0],[0,9],[7],[7,9],[0,7],[0,7,9],[3],[3,9],[0,3],[0,3,9],[3,7],[3,7,9],[0,3,7],[0,3,7,9],[5],[5,9],[0,5],[0,5,9],[5,7],[5,7,9],[0,5,7],[0,5,7,9],[3,5],[3,5,9],[0,3,5],[0,3,5,9],[3,5,7],[3,5,7,9],[0,3,5,7],[0,3,5,7,9]]
	//[[] [9] [0] [9 0] [7] [9 7] [0 7] [9 0 7] [3] [9 3] [0 3] [9 0 3] [7 3] [9 7 3] [0 7 3] [9 0 7 5] [5] [9 5] [0 5] [9 0 5] [7 5] [9 7 5] [0 7 5] [9 0 7 5] [3 5] [9 3 5] [0 3 5] [9 0 3 5] [7 3 5] [9 7 3 5] [0 7 3 5] [9 0 7 5 5]]
	//nums := []int{1, 2, 3}
	nums := []int{9, 0, 7, 3, 5}
	i := subsets(nums)
	fmt.Println(i) // [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}
/*
递归：
位运算：
dp：
回溯：
*/
//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	// 回溯
	subs := make([][]int, 0)
	dfs(0, []int{}, &subs, nums)
	return subs

	// dp
	//dp := make([][]int, 1, 1<<len(nums))
	//dp[0] = []int{}
	//for i := 0; i < len(nums); i++ {
	//	n := len(dp)
	//	for j := 0; j < n; j++ {
	//		subPlus := make([]int, 0)
	//		subPlus = append(subPlus, dp[j]...)
	//		dp = append(dp, append(subPlus, nums[i]))
	//		// 刚好是在 8 这个位置，地址相同
	//		//dp = append(dp, append(dp[j], nums[i]))
	//	}
	//}
	//return dp

	// 位运算
	//n := len(nums)
	//count := 1 << n
	//subs := make([][]int, count)
	//for i := 0; i < count; i++ {
	//	sub := make([]int, 0)
	//	for j, idx := i, 0; j > 0 && idx < n; j, idx = j>>1, idx+1 {
	//		if j&1 == 1 {
	//			sub = append(sub, nums[idx])
	//		}
	//	}
	//	subs[i] = sub
	//}
	//return subs

	// 递归：2
	//subs := make([][]int, 1)
	//subs[0] = []int{}
	//subsetsRecursion(&subs, 0, nums)
	//return subs

	// 递归：但是题目要求结果的顺序
	//subs := make([][]int, 1)
	//subs[0] = []int{}
	//subsetsRecursion(&subs, subs[0], nums)
	//return subs
}

func dfs(i int, ints []int, subs *[][]int, nums []int) {
	sub := make([]int, len(ints))
	copy(sub, ints)
	*subs = append(*subs, sub)
	//*subs = append(*subs, ints)
	for ; i < len(nums); i++ {
		ints = append(ints, nums[i])
		dfs(i+1, ints, subs, nums)
		ints = ints[:len(ints)-1]
	}
}

//func subsetsRecursion(subs *[][]int, sub []int, nums []int) {
func subsetsRecursion(subs *[][]int, i int, nums []int) {
	if i == len(nums) {
		return
	}
	n := len(*subs)
	for j := 0; j < n; j++ {
		//*subs = append(*subs, append((*subs)[j], nums[i]))
		subPlus := make([]int, 0)
		subPlus = append(subPlus, (*subs)[j]...)
		*subs = append(*subs, append(subPlus, nums[i]))
	}
	subsetsRecursion(subs, i+1, nums)

	//if len(nums) == 0 {
	//	return
	//}
	//for i := 0; i < len(nums); i++ {
	//	subPlus := append(sub, nums[i])
	//	*subs = append(*subs, subPlus)
	//	subsetsRecursion(subs, subPlus, nums[i+1:])
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
