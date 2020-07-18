//给你一个整数数组 nums，请你将该数组升序排列。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
// 
//
// 示例 2： 
//
// 输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 50000 
// -50000 <= nums[i] <= 50000 
// 
//
package main

import "fmt"

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	array := sortArray(nums)
	fmt.Println(array)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	// 堆排序
	n := len(nums)
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, 0, i)
	}
	return nums
}

func heapify(arr []int, i int, n int) {
	for idx := i<<1 + 1; idx < n; idx = idx<<1 + 1 {
		if idx+1 < n && arr[idx+1] > arr[idx] {
			idx++
		}
		if arr[i] >= arr[idx] {
			break
		}
		arr[i], arr[idx], i = arr[idx], arr[i], idx
	}
}

//leetcode submit region end(Prohibit modification and deletion)
