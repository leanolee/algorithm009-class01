//假设按照升序排序的数组在预先未知的某个点上进行了旋转。 
//
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 
//
// 请找出其中最小的元素。 
//
// 你可以假设数组中不存在重复元素。 
//
// 示例 1: 
//
// 输入: [3,4,5,1,2]
//输出: 1 
//
// 示例 2: 
//
// 输入: [4,5,6,7,0,1,2]
//输出: 0 
// Related Topics 数组 二分查找
package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{4, 5, 6, 2}
	//nums := []int{4}
	//nums := []int{1, 2, 3} // 这也算旋转数组？
	min := findMin(nums)
	fmt.Println(min)
}

/*
思路：可以参考：BinarySearch.go
迭代：O(n)，当有重复元素时，只能使用迭代
二分：旋转数组寻找最小值，使用二分的前提条件是，数组不包含重复元素
	此题认为 [1,2,3] 是旋转数组
*/
//leetcode submit region begin(Prohibit modification and deletion)
func findMin(nums []int) int {
	// 二分：精简
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]

	// 二分：官方
	//l, r := 0, len(nums)-1
	//if nums[l] <= nums[r] {
	//	return nums[0]
	//}
	//for l < r {
	//	mid := (l + r) >> 1
	//	if mid > 0 && nums[mid] < nums[mid-1] {
	//		return nums[mid]
	//	} else if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
	//		return nums[mid+1]
	//	}
	//	if nums[mid] > nums[l] {
	//		l = mid + 1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return nums[l]

	// 二分：判断首尾
	//l, r := 0, len(nums)-1
	//for l < r {
	//	mid := (l + r) >> 1
	//	if nums[l] < nums[r] {
	//		return nums[l]
	//	}
	//	if nums[mid] >= nums[l] {
	//		l = mid + 1
	//	} else {
	//		r = mid
	//	}
	//}
	//return nums[r]

	// 遍历：O(n)
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] < nums[i-1] {
	//		return nums[i]
	//	}
	//}
	//return nums[0]

	// 寻找最大值：这种写法是，寻找最大的数
	//l, r := 0, len(nums)
	//if r == 1 {
	//	return nums[0]
	//}
	//for l < r {
	//	mid := (l + r) >> 1
	//	if nums[mid] > nums[l] {
	//		l = mid
	//	} else {
	//		r = mid
	//	}
	//}
	//return nums[l]
}

//leetcode submit region end(Prohibit modification and deletion)
