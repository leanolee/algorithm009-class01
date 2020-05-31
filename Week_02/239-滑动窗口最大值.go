//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。 
//
//
// 返回滑动窗口中的最大值。 
//
// 
//
// 进阶： 
//
// 你能在线性时间复杂度内解决此题吗？ 
//
// 
//
// 示例: 
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7] 
//解释: 
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10^5 
// -10^4 <= nums[i] <= 10^4 
// 1 <= k <= nums.length 
// 
// Related Topics 堆 Sliding Window
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//nums := []int{1}
	k := 3
	window := maxSlidingWindow(nums, k)
	fmt.Println(window)
}

/*
暴力法：
变量：O(nk),O(1)
	https://leetcode.com/problems/sliding-window-maximum/discuss/437576/Java-1ms-solution-using-only-array-and-easy-to-understand
	思路：
		1.用一个变量 maxIdx 记录 i ~ i+k-1 范围内最大数的 index
		2.每次循环，先判断，新进来的元素的值，取 maxIdx 和 i 对应的值中最大的值为（新的） maxIdx
		3.如果 maxIdx 比当前的索引位置 i 大，说明 maxIdx 还在有效的范围内（可以被使用）
		4.如果 maxIdx 比当前的索引位置 i 小，说明 maxIdx 需要被淘汰
			此时去遍历 i ~ i+k-1 之间最大的数，并记录新的 maxIdx
		5.经过上述判断后，将 maxIdx 位置的值添加到结果数组中
dp：
deque：
	假设 k = 3
	第一次进循环，只会 q.offer(i)

	deque：O(n),O(n)
	双端队列：所有的 sliding window（滑动窗口）的题目，都用队列（双端队列）去处理
	1.创建 k 长的deque，新进元素从尾到头依次和 deque 中元素比较
		比较 2n 次：分别和前一个、后一个元素比较一次
		小 则直接放入
		大 则删掉前面的元素
		根据删掉了几个元素，进行截取
	2.从 deque 中获取 max，并判断是否要去掉超出 k 范围的 左端 元素
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	// 变量
	n := len(nums)
	window := make([]int, n-k+1)
	maxIdx := -1
	for i := 0; i <= n-k; i++ {
		if maxIdx != -1 && nums[i+k-1] >= nums[maxIdx] {
			maxIdx = i + k - 1
		} else if i > maxIdx {
			maxIdx = i
			for p := i + 1; p < i+k; p++ {
				if nums[p] >= nums[maxIdx] {
					maxIdx = p
				}
			}
		}
		window[i] = nums[maxIdx]
	}
	return window

	// deque
	//n := len(nums)
	//window := make([]int, n-k+1)
	//deque := make([]int, 0)
	//for i := 0; i < n; i++ {
	//	delCount := 0
	//	for j := len(deque) - 1; j >= 0 && nums[i] > nums[deque[j]]; j-- {
	//		delCount++
	//	}
	//	deque = deque[:len(deque)-delCount]
	//	deque = append(deque, i)
	//	if i-deque[0] == k {
	//		deque = deque[1:]
	//	}
	//	if i+1 >= k {
	//		window[i+1-k] = nums[deque[0]]
	//	}
	//}
	//return window

	// dp
	//n := len(nums)
	//window := make([]int, n-k+1)
	//left := make([]int, n)
	//right := make([]int, n)
	//maxLeft, maxRight := nums[0], nums[n-1]
	//for i, j := 0, n-1; i < n; {
	//	if i%k == 0 {
	//		maxLeft = nums[i]
	//	} else {
	//		maxLeft = MyTools.Max(nums[i], maxLeft)
	//	}
	//	if j%k == k-1 {
	//		maxRight = nums[j]
	//	} else {
	//		maxRight = MyTools.Max(nums[j], maxRight)
	//	}
	//	left[i] = maxLeft
	//	i++
	//	right[j] = maxRight
	//	j--
	//}
	//for i := 0; i < n-k+1; i++ {
	//	window[i] = MyTools.Max(left[i+k-1], right[i])
	//}
	//return window
}

//leetcode submit region end(Prohibit modification and deletion)
