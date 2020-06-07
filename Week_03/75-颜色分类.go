//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。 
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。 
//
// 注意: 
//不能使用代码库中的排序函数来解决这道题。 
//
// 示例: 
//
// 输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2] 
//
// 进阶： 
//
// 
// 一个直观的解决方案是使用计数排序的两趟扫描算法。 
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。 
// 你能想出一个仅使用常数空间的一趟扫描算法吗？ 
// 
// Related Topics 排序 数组 双指针
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	//nums := []int{1, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}

/*

 */
//leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	// 双指针
	l, r := 0, len(nums)-1
	for i := 0; i <= r; {
		switch nums[i] {
		case 0:
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		case 2:
			nums[r], nums[i] = nums[i], nums[r]
			r--
		case 1:
			i++ // 0 1 是已经遍历过的数，才需要 i++
		}
	}

	// 计数法
	//couter := [3]int{}
	//for _, num := range nums {
	//	couter[num]++
	//}
	//for i := 0; i < len(nums); i++ {
	//	if i < couter[0] {
	//		nums[i] = 0
	//	} else if i >= len(nums)-couter[2] {
	//		nums[i] = 2
	//	} else {
	//		nums[i] = 1
	//	}
	//}

	// sort：这里使用快排
	//sortQ75(nums, 0, len(nums)-1)
}

func sortQ75(nums []int, begin int, end int) {
	if begin >= end {
		return
	}
	pivot := partition75(nums, begin, end)
	sortQ75(nums, begin, pivot-1)
	sortQ75(nums, pivot+1, end)
}

func partition75(nums []int, begin int, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	return counter
}

//leetcode submit region end(Prohibit modification and deletion)
