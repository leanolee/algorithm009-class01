package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	sort := binarySearchHalfSort(nums, 2)
	fmt.Println(sort)
}
/*
二分查找一个旋转数组：
*/
func binarySearchHalfSort(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[r] {
			if nums[mid] > target && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		//if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
		//	right = mid - 1
		//} else {
		//	left = mid + 1
		//}
	}
	return -1
}
