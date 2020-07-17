//给定一个无序的整数数组，找到其中最长上升子序列的长度。 
//
// 示例: 
//
// 输入: [10,9,2,5,3,7,101,18]
//输出: 4 
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。 
//
// 说明: 
//
// 
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。 
// 你算法的时间复杂度应该为 O(n2) 。 
// 
//
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗? 
// Related Topics 二分查找 动态规划
package main

import "fmt"

func main() {
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{2, 2}
	//nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	nums := []int{}
	lis := lengthOfLIS(nums)
	fmt.Println(lis)
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	// 二分：
	//cache, l, r, n := make([]int, 0), 0, 0, 0
	//for _, v := range nums {
	//	if n == 0 || v > cache[n-1] {
	//		cache, n = append(cache, v), n+1
	//		continue
	//	}
	//	l, r = 0, n
	//	for l < r {
	//		mid := (l + r) >> 1
	//		if v > cache[mid] {
	//			l = mid + 1
	//		} else {
	//			r = mid
	//		}
	//	}
	//	cache[r] = v
	//}
	//return len(cache)

	// DP:
	n := len(nums)
	dp, max := make([]int, n), 0
	for i := 0; i < n; i++ {
		currMax := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				currMax = getMax(currMax, dp[j])
			}
		}
		dp[i] = currMax + 1
	}
	for i := 0; i < n; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
