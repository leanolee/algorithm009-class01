//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。 
//
// 示例 1: 
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
// 
//
// 示例 2: 
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4 
//
// 说明: 
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。 
// Related Topics 堆 分治算法
package main

import "fmt"

func main() {
	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	nums := []int{13, 54, 32, 654, 889, 44, 0, 34, 987, 45, 765, 23, 76, 3, 2, 6}
	k := 4
	largest := findKthLargest(nums, k)
	fmt.Println(largest)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	// 堆
	//heapify := func(arr []int, i, n int) {
	//	for idx := i<<1 + 1; idx < n; idx = idx<<1 + 1 {
	//		if idx+1 < n && arr[idx+1] > arr[idx] {
	//			idx++
	//		}
	//		if arr[idx] > arr[i] {
	//			arr[idx], arr[i], i = arr[i], arr[idx], idx
	//		}
	//	}
	//}
	//n := len(nums)
	//for i := n>>1 - 1; i >= 0; i-- {
	//	heapify(nums, i, n)
	//}
	//idx := n - k
	//for i := n - 1; i > idx; i-- {
	//	nums[0], nums[i] = nums[i], nums[0]
	//	heapify(nums, 0, i)
	//}
	//return nums[0]

	// 快排
	//patition := func(nums []int, l, r int) int {
	//	pivot, counter := r, l
	//	for i := l; i < r; i++ {
	//		if nums[i] < nums[pivot] {
	//			nums[counter], nums[i] = nums[i], nums[counter]
	//			counter++
	//		}
	//	}
	//	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	//	return counter
	//}
	//var qSort func(nums []int, l, r int, k int) int
	//qSort = func(nums []int, l, r int, k int) int {
	//	if l >= r {
	//		return nums[r]
	//	}
	//	pivot := patition(nums, l, r)
	//	switch {
	//	case pivot == k:
	//		return nums[pivot]
	//	case pivot < k:
	//		return qSort(nums, pivot+1, r, k)
	//	case pivot > k:
	//		return qSort(nums, l, pivot-1, k)
	//	}
	//	return 0
	//}
	//return qSort(nums, 0, len(nums)-1, len(nums)-k)

	// 半排序：如冒泡排序：O(nk)

	// 全排序：复习堆
	n := len(nums)
	fmt.Println(n, n>>1-1)
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}
	fmt.Println(nums)
	return nums[n-k]
}

func heapify(arr []int, i int, n int) {
	for idx := i<<1 + 1; idx < n; idx = idx<<1 + 1 {
		if idx+1 < n && arr[idx+1] > arr[idx] {
			idx++
		}
		if arr[idx] > arr[i] { // else break
			arr[idx], arr[i], i = arr[i], arr[idx], idx
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
