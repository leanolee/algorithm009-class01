//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。 
//
// 
//
// 说明: 
//
// 
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。 
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。 
// 
//
// 
//
// 示例: 
//
// 输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6] 
// Related Topics 数组 双指针
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	//nums1 := []int{0}
	//nums2 := []int{2}
	//m, n := 0, 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	//copy(nums1[2:], nums1[1:])
	//fmt.Println("num1:", nums1)
}
/*
双指针（从前往后）：O(n),O(1)
	思路：
	1.循环遍历，i 记录当前要操作的索引
	2.j、k，分别对应 num1、num2 当前要操作的索引。但其实 num1 的位置（不一定是索引了）其实是由 i 记录
	3.k == n，说明 短 数组已经操作完，可直接返回
	4.三步判断：
		nums1[i] > nums2[k]：长 数组的当前值大
			则往后 copy 一位，nums1[i] = nums2[k]
			k++
		nums1[i] <= nums2[k]：
			j < m：说明 num1 还没处理完，此时不需要任何操作。j++
			j==m：说明 num1 已处理完，只用拷贝 num2 的值就行。k++
双指针：O(n),O(1)
	从后往前，更简洁高效
	1.利用参数 m、n，记录剩余需要操作的两个数组的个数
	2.m < 1：说明 num1 已处理完，只用拷贝 num2 的值就行。k++
	3.n < 1 || nums1[m-1] > nums2[n-1]：说明 num2 已处理完，或 num1 值更大，拷贝 num1的值
	4.其余：拷贝 num2 值
*/
//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 双指针
	//for i := m + n - 1; i >= 0; i-- {
	//	if m < 1 {
	//		nums1[i] = nums2[n-1]
	//		n--
	//	} else if n < 1 || nums1[m-1] > nums2[n-1] {
	//		nums1[i] = nums1[m-1]
	//		m--
	//	} else {
	//		nums1[i] = nums2[n-1]
	//		n--
	//	}
	//}

	// 双指针：从前
	for i, j, k := 0, 0, 0; i < len(nums1); i++ {
		if k == n {
			break
		}
		if nums1[i] > nums2[k] {
			copy(nums1[i+1:i+1+m-j], nums1[i:i+m-j])
			nums1[i] = nums2[k]
			k++
		} else if j < m {
			j++
			continue
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
