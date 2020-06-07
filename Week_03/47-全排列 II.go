//给定一个可包含重复数字的序列，返回所有不重复的全排列。 
//
// 示例: 
//
// 输入: [1,1,2]
//输出:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//] 
// Related Topics 回溯算法
package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{0, 1, 0, 0, 9} // 20
	//nums := []int{2, 2, 1, 1} // 6
	nums := []int{-1, 2, 0, -1, 1, 0, 1} // 630
	unique := permuteUnique(nums)
	fmt.Println(unique, len(unique))
}

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	// 递归：回溯+hash
	//unique := make([][]int, 0)
	//permuteUniqueDFS(nums, &unique, 0)
	//return unique

	// 递归：回溯+排序
	sort.Ints(nums)
	unique := make([][]int, 0)
	permuteUniqueDFS_(nums, &unique, 0)
	return unique
}

func permuteUniqueDFS(nums []int, unique *[][]int, idx int) {
	if idx == len(nums)-1 {
		*unique = append(*unique, append([]int{}, nums...))
		return
	}
	memo := make(map[int]bool)
	for i := idx; i < len(nums); i++ {
		if _, ok := memo[nums[i]]; ok {
			continue
		}
		nums[i], nums[idx] = nums[idx], nums[i]
		permuteUniqueDFS(nums, unique, idx+1)
		nums[i], nums[idx] = nums[idx], nums[i]
		memo[nums[i]] = true
	}
}

func permuteUniqueDFS_(nums []int, unique *[][]int, idx int) {
	if idx == len(nums)-1 {
		*unique = append(*unique, append([]int{}, nums...))
		return
	}
	for i := idx; i < len(nums); i++ {
		//if i > idx && (nums[i] == nums[idx] || nums[i] == nums[i-1]) {	// 错误
		if !checked(nums, idx, i) {	// 使用 leetcode 的 used 方法，不如直接使用 hash
			continue
		}
		nums[i], nums[idx] = nums[idx], nums[i]
		permuteUniqueDFS_(nums, unique, idx+1)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}

func checked(nums []int, idx int, i int) bool {
	for j := idx; j < i; j++ {
		if nums[j] == nums[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
