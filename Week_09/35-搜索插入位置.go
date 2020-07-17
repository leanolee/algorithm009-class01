//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。 
//
// 你可以假设数组中无重复元素。 
//
// 示例 1: 
//
// 输入: [1,3,5,6], 5
//输出: 2
// 
//
// 示例 2: 
//
// 输入: [1,3,5,6], 2
//输出: 1
// 
//
// 示例 3: 
//
// 输入: [1,3,5,6], 7
//输出: 4
// 
//
// 示例 4: 
//
// 输入: [1,3,5,6], 0
//输出: 0
// 
// Related Topics 数组 二分查找
package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 8
	i := searchInsert(nums, target)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	// 二分
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
