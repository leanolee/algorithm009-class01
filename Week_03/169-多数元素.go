//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 
//
// 示例 1: 
//
// 输入: [3,2,3]
//输出: 3 
//
// 示例 2: 
//
// 输入: [2,2,1,1,1,2,2]
//输出: 2
// 
// Related Topics 位运算 数组 分治算法
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2, 2, 2, 1, 1, 1, 2, 2}
	element := majorityElement(nums)
	fmt.Println(element)
}

/*
暴力法：O(n),O(n)
	hash
排序：
	下标为 n/2 的数，就是众数
随机化：
分治：
	相同，则返回
	不相同，统计后返回
位运算：
*/
//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// 摩尔投票法:Boyer-Moore 投票
	majority, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			count++
		} else {
			if count == 0 {
				majority = nums[i]
				count++
			} else {
				count--
			}
		}
	}
	return majority

	// 分治
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	//mid := len(nums) >> 1
	//left := majorityElement(nums[:mid])
	//right := majorityElement(nums[mid:])
	//return mergerMajority(left, right, nums)

	// 排序：这里使用快排
	//quickSort169(nums, 0, len(nums)-1)
	//return nums[len(nums)/2]

	// 暴力法
	//memo := make(map[int]int)
	//for _, n := range nums {
	//	memo[n]++
	//}
	//num, max := nums[0], 1
	//for k, v := range memo {
	//	if v > max {
	//		max = v
	//		num = k
	//	}
	//}
	//return num
}

func quickSort169(nums []int, begin int, end int) {
	if begin >= end {
		return
	}
	pivot := partition169(nums, begin, end)
	quickSort169(nums, begin, pivot-1)
	quickSort169(nums, pivot+1, end)
}

func partition169(nums []int, begin int, end int) int {
	pivot, couter := end, begin
	for i := begin; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[couter], nums[i] = nums[i], nums[couter]
			couter++
		}
	}
	nums[couter], nums[pivot] = nums[pivot], nums[couter]
	return couter
}

func mergerMajority(left int, right int, nums []int) int {
	if left == right {
		return left
	}
	lCount, rCount := 0, 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case left:
			lCount++
		case right:
			rCount++
		}
	}
	if lCount > rCount {
		return left
	}
	return right
}

//leetcode submit region end(Prohibit modification and deletion)
