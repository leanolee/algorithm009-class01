//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例： 
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sum := threeSum(nums)
	fmt.Println(sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	// 双指针
	sort.Ints(nums)
	n, iN := len(nums)-1, len(nums)-3
	ans := make([][]int, 0)
	for i := 0; i <= iN; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, n; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
