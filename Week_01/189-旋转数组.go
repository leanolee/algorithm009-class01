//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。 
//
// 示例 1: 
//
// 输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
// 
//
// 示例 2: 
//
// 输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释: 
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100] 
//
// 说明: 
//
// 
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。 
// 要求使用空间复杂度为 O(1) 的 原地 算法。 
// 
// Related Topics 数组
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//nums := []int{1}
	//nums := []int{-1, -100, 3, 99}
	//nums := []int{1, 2, 3, 4, 5, 6}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
/*
暴力法：O(n*k),O(1)
	循环 k 次，每一次挪动一位
迭代：O(n),O(k)
	不推荐的笨办法
	1.创建 k 长度数组，分别保存每次循环时，每次 k 跳的元素
	2.循环数组
		for i := 0; i < n+1; i += k {
			for j := 0; j < k; j++ {
	3.补刀：由于这种办法，最后会有剩余元素，所以需要补全
		for i := 0; i < n%k; i++ {
			nums[k-n%k+i] = memo[i]
		}
迭代+数组：O(n),O(n)
	1.新建一个原数组的拷贝
	2.两个数组，错开 k 位拷贝
迭代+数组优化：O(n),O(k)
	1.创建 k 长数组，拷贝原数组的前 k 位
	2.进行 k 跳的循环，利用 k 长数组进行值的交换
反转：O(n),O(1)
	1.对整个数组进行反转
	2.对前 k 位进行反转，对剩余部分进行反转
*/
//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	//反转
	//n := len(nums)
	//k %= n
	//reverse(nums, 0, n-1)
	//reverse(nums, 0, k-1)
	//reverse(nums, k, n-1)

	// 纯移动：优化，失败。
	//n := len(nums)
	//k %= n
	//for i := 0; i < k; i++ {
	//	temp := nums[i]
	//	start := i
	//	for j := k + i; j < n+k; j += k {
	//		if j >= n {
	//			j -= n
	//		}
	//		temp, nums[j] = nums[j], temp
	//		fmt.Println(j)
	//		if j == start {
	//			if n%k == 0 {
	//				break
	//			} else {
	//				return
	//			}
	//		}
	//	}
	//}

	// 纯移动：失败。原数组上操作。如当n=6，k=4时，j的取值在 0、4、8->2、6->0...之间如此循环
	//n := len(nums)
	//k %= n
	//temp := nums[0]
	//off := n % k
	//if off == 0 {
	//	for i := 0; i < k; i++ {
	//		for j := i; j < n+k; j += k {
	//			index := j
	//			if index >= n {
	//				index -= n
	//			}
	//			temp, nums[index] = nums[index], temp
	//		}
	//	}
	//} else {
	//	for j := k; j < n+k; j += k {
	//		if j >= n {
	//			j -= n
	//		}
	//		temp, nums[j] = nums[j], temp
	//		if j == 0 {
	//			break
	//		}
	//	}
	//}

	// 两数组优化：
	n := len(nums)
	k %= n
	memo := make([]int, k)
	for i := 0; i < k; i++ {
		memo[i] = nums[i]
	}
	for i := 0; i < k; i++ {
		for j := k + i; j < n+k; j += k {
			if j >= n {
				j -= n
			}
			memo[i], nums[j] = nums[j], memo[i]
			if j < k {
				break
			}
		}
	}

	// 两数组暴力法：暴力法不是一个个的旋转吗？所以这个是使用两个数组
	//n := len(nums)
	//temp := make([]int, n)
	//for i, num := range nums {
	//	temp[i] = num
	//}
	//for i := 0; i < n; i++ {
	//	index := i + k
	//	for index >= n {
	//		index -= n
	//	}
	//	nums[index] = temp[i]
	//}

	// memo：
	//n := len(nums)
	//if n < 2 {
	//	return
	//}
	//k %= n
	//memo := make([]int, k)
	////for i := 0; i < n+k; i += k {	// 如何合并呢
	//for i := 0; i < n+1; i += k {
	//	for j := 0; j < k; j++ {
	//		nIdx := i + j
	//		if nIdx >= n {
	//			nIdx -= n
	//		}
	//		memo[j], nums[nIdx] = nums[nIdx], memo[j]
	//	}
	//}
	//for i := 0; i < n%k; i++ {
	//	nums[k-n%k+i] = memo[i]
	//}

	// 真正的暴力法
	//n := len(nums)
	//k %= n
	//for i := 0; i < k; i++ {
	//	temp := nums[n-1]
	//	for j := 0; j < n; j++ {
	//		temp, nums[j] = nums[j], temp
	//	}
	//}
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
