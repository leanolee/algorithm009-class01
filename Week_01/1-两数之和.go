//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。 
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。 
//
// 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := twoSum(nums, target)
	fmt.Println(sum)
}

/*
暴力法：O(n^2),O(1)
	1.一层循环：固定第一个数
	2.二层循环：遍历数组，分别和第一个数相加
memo：O(n),O(n)
	1.创建memo
	2.遍历数组：分为3部分理解
		1.已遍历的数组，将余放进memo中
		2.未遍历的数组
		3.当下的数，如果在memo中有该值，则找到
 */
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	// memo
	memo := make(map[int]int)
	for i, num := range nums {
		if index, ok := memo[num]; ok {
			return []int{index, i}
		}
		memo[target-num] = i
	}
	return nil

	// 暴力法
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return nil
}

//leetcode submit region end(Prohibit modification and deletion)
