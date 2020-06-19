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

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := twoSum(nums, target)
	fmt.Println(sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	// hash
	memo := make(map[int]int)
	for i, v := range nums {
		if idx, ok := memo[v]; ok {
			return []int{idx, i}
		}
		memo[target-v] = i
	}
	return nil

	// n数之和
	//sort.Ints(nums)
	//sums := make([][]int, 0)
	//n := 2
	//nSum(nums, &sums, nil, n, target) // n数之和
	//return sums[0]
}

func nSum(nums []int, sums *[][]int, path []int, n int, target int) {
	if n < 2 || len(nums) < n || nums[0]*n > target {
		return
	}
	if n == 2 {
		i, j := 0, len(nums)-1
		for i < j {
			if nums[i] > target {
				break
			}
			sum := nums[i] + nums[j]
			switch {
			case sum == target:
				*sums = append(*sums, append(append([]int{}, path...), i, j))
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			case sum < target:
				i++
			case sum > target:
				j--
			}
		}
	} else {
		for i := 0; i < len(nums)-n+1; i++ { //i < len(nums)-n+1
			if i == 0 || nums[i] != nums[i-1] {
				nSum(nums[i+1:], sums, append(append([]int{}, path...), i), n-1, target-nums[i]) // nums[i+1:]
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
