//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。 
//
// 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。 
//
// 说明: 
//
// 
// 返回的下标值（index1 和 index2）不是从零开始的。 
// 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。 
// 
//
// 示例: 
//
// 输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。 
// Related Topics 数组 双指针 二分查找
package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 26
	sum := twoSum(numbers, target)
	fmt.Println(sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	// 双指针
	for i, j := 0, len(numbers)-1; i < j; {
		switch {
		case numbers[i]+numbers[j] == target:
			return []int{i + 1, j + 1}
		case numbers[i]+numbers[j] < target:
			i++
		case numbers[i]+numbers[j] > target:
			j--
		}
	}
	return nil

	// 二分：O(n log n)，因为有序，所以可以二分找差值
}

//leetcode submit region end(Prohibit modification and deletion)
