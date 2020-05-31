//给定两个数组，编写一个函数来计算它们的交集。 
//
// 示例 1: 
//
// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2,2]
// 
//
// 示例 2: 
//
// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9] 
//
// 说明： 
//
// 
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。 
// 我们可以不考虑输出结果的顺序。 
// 
//
// 进阶: 
//
// 
// 如果给定的数组已经排好序呢？你将如何优化你的算法？ 
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？ 
// 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？ 
// 
// Related Topics 排序 哈希表 双指针 二分查找
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	ints := intersect(nums1, nums2)
	fmt.Println(ints)
}

/*
双指针：O(nlogn),O(1)
hash：O(m)/O(n),O(m+n)
hash优化：O(m)/O(n),O(m+n)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	// hash优化
	ints := make([]int, 0)
	memo := make(map[int]int)
	for _, num := range nums1 {
		memo[num] = memo[num] + 1
	}
	for _, num := range nums2 {
		v := memo[num]
		if v > 0 {
			ints = append(ints, num)
			memo[num] = v - 1
		}
	}
	return ints

	// hash
	//ints := make([]int, 0)
	//memo := make(map[int]int)
	//for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
	//	if i < len(nums1) {
	//		v := memo[nums1[i]]
	//		if v < 0 {
	//			ints = append(ints, nums1[i])
	//		}
	//		memo[nums1[i]] = v + 1
	//		i++
	//	}
	//	if j < len(nums2) {
	//		v := memo[nums2[j]]
	//		if v > 0 {
	//			ints = append(ints, nums2[j])
	//		}
	//		memo[nums2[j]] = v - 1
	//		j++
	//	}
	//}
	//return ints

	// 双指针
	//sort.Ints(nums1)
	//sort.Ints(nums2)
	//ints := make([]int, 0)
	//for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
	//	switch {
	//	case nums1[i] > nums2[j]:
	//		j++
	//	case nums1[i] < nums2[j]:
	//		i++
	//	case nums1[i] == nums2[j]:
	//		ints = append(ints, nums1[i])
	//		i++
	//		j++
	//	}
	//}
	//return ints
}

//leetcode submit region end(Prohibit modification and deletion)
