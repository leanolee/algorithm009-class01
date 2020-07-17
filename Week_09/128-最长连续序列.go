//给定一个未排序的整数数组，找出最长连续序列的长度。 
//
// 要求算法的时间复杂度为 O(n)。 
//
// 示例: 
//
// 输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。 
// Related Topics 并查集 数组
package main

import (
	"fmt"
)

func main() {
	//nums := []int{100, 4, 200, 1, 3, 2}
	nums := []int{-1, 1, 0}
	subsequence := longestConsecutive(nums)
	fmt.Println(subsequence)
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	// hash
	memo := make(map[int]bool)
	for _, v := range nums {
		memo[v] = true
	}
	count, curr := 0, 1
	for k, _ := range memo {
		for v := k-1; memo[v]; v-- {
			curr++
			delete(memo, v)
		}
		for v := k+1; memo[v]; v++ {
			curr++
			delete(memo, v)
		}
		if curr > count {
			count = curr
		}
		curr = 1
	}
	return count

	// 并查集：https://leetcode-cn.com/problems/longest-consecutive-sequence/solution/128-zui-chang-lian-xu-xu-lie-liang-chong-ha-xi-fan/


	// 全排序
	//if len(nums) < 2 {
	//	return len(nums)
	//}
	//sort.Ints(nums)
	//count, curr := 1, 1
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] == nums[i-1] {
	//		continue
	//	}
	//	if nums[i] == nums[i-1]+1 {
	//		curr++
	//		if curr > count {
	//			count = curr
	//		}
	//	} else {
	//		curr = 1
	//	}
	//}
	//return count
}

//leetcode submit region end(Prohibit modification and deletion)
