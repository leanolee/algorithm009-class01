package main

import "fmt"

func main() {
	//nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, -4, -3, -2, -1, 0, 1, 2, 3}
	//nums := []int{10, 11, 12, -4, -3}
	//nums := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, -4, -4, -3, -2, -1, 0, 1, 2, 3}
	//nums := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 0, 0, 0, 1, 1, 2, 2, 3}
	nums := []int{3, 0, 3, 3, 3, 3}
	sort := idxOfHalfSort(nums)
	fmt.Println("idx:", sort)
	fmt.Println("val:", nums[sort])
}

/*
需求：一个旋转数组，找到它旋转的起始位置
思路：
	即找到索引 i，nums[i] < nums[i-1]
方法一：假设没有重复元素
	遍历：O(n)
		找到小于前一个数的数的索引
	二分：O(logn)
		if 首元素 < 尾元素：放弃
		else if nums[begin] >= nums[end] {
			// 继续二分
		}
二分：
	l, r := 0, len(nums)-1
	if nums[mid] > nums[r] {
		l = mid + 1
	} else {
		r = mid - 1
	}
	存在问题：如果 mid 恰好是要找的元素，则被跳过了
	解决办法：每次先判断 nums[mid] < nums[mid-1]
方法二：假设有重复元素（但至少有2个不同的元素）
	重点：有3种情况需要关注
		1.mid 是旋转起始位：
		2.mid 的右边那个元素是旋转起始位：
		3.mid 的左边那个元素是旋转起始位：这种情况不需要处理，因为它被包含在了 nums[mid] < nums[l] 的情况中
	打脸：[]int{3, 0, 3, 3, 3, 3}
		这种情况没法通过二分来确定旋转起始位
结论：通过二分来确定一个旋转数组的起始位
	前提条件：数组没有重复元素
*/
func idxOfHalfSort(nums []int) int {
	// 方法二：假设有重复元素（但至少有2个不同的元素）
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if mid > 0 && nums[mid] < nums[mid-1] { // 1.mid 是旋转起始位
			return mid
		}
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] { // 2.防止 mid 的右边那个元素是旋转起始位
			return mid + 1
		}
		if nums[mid] >= nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l

	// 方法一：假设没有重复元素
	//l, r := 0, len(nums)-1
	//for l < r {
	//	if nums[l] < nums[r] {
	//		return nums[l]
	//	}
	//	if r-l == 1 {
	//		return nums[r]
	//	}
	//	mid := (l + r + 1) >> 1
	//	if nums[l] < nums[mid] {
	//		l = mid
	//	} else {
	//		r = mid
	//	}
	//}
	//return nums[l]
}
