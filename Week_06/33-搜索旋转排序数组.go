//假设按照升序排序的数组在预先未知的某个点上进行了旋转。 
//
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 
//
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 
//
// 你可以假设数组中不存在重复的元素。 
//
// 你的算法时间复杂度必须是 O(log n) 级别。 
//
// 示例 1: 
//
// 输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
// 
//
// 示例 2: 
//
// 输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1 
// Related Topics 数组 二分查找
package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{4, 5, 6, 7, 8, 9, 10, 11}
	//nums := []int{1, 3}
	target := 4
	i := search(nums, target)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[l] && (nums[mid] < target || nums[mid] > target && target < nums[l]) || nums[mid] < nums[l] && nums[mid] < target && nums[l] > target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		fmt.Println(l, r, mid)
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
