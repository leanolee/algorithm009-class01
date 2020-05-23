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
	"./MyTools"
	"fmt"
)

func main() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	window := maxSlidingWindow(nums, k)
	fmt.Println(window)
}

/*
暴力法：O(n*k),O(n)
	1.一层循环：固定当前 k 个数的起始位置
	2.二层循环：遍历 k 个数，寻找最大值
deque：O(n),O(n)
	双端队列：所有的 sliding window（滑动窗口）的题目，都用队列（双端队列）去处理
	1.创建 k 长的deque，新进元素从尾到头依次和 deque 中元素比较
		比较 2n 次：分别和前一个、后一个元素比较一次
		小 则直接放入
		大 则删掉前面的元素
		根据删掉了几个元素，进行截取
	2.从 deque 中获取 max，并判断是否要去掉超出 k 范围的 左端 元素
动态规划：O(n),O(n)
	与 leetcode-42有异曲同工之妙
	1.创建两个数组 left、right，分别记录从 左->右、右->左 的dp后的元素
	2.每 k 个元素一组，尾巴上剩余的元素为一组
		left、right记录对应位置上，递推的最大值
		每过 k 个，重置递推
	3.遍历 left、right，取 每k个元素中 max(left[i+k-1], right[i])
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	// dp
	n := len(nums)
	memo := make([][2]int, n)
	dp := make([]int, n-k+1)
	maxLeft, maxRight := nums[0], nums[n-1]
	for i, j := 0, n-1; i < n; {
		if i%k == 0 {
			maxLeft = nums[i]
		} else {
			maxLeft = MyTools.Max(nums[i], maxLeft)
		}
		memo[i][0] = maxLeft
		i++
		if j%k == k-1 {
			maxRight = nums[j]
		} else {
			maxRight = MyTools.Max(nums[j], maxRight)
		}
		memo[j][1] = maxRight
		j--
	}
	for i := 0; i < n-k+1; i++ {
		dp[i] = MyTools.Max(memo[i][1], memo[i+k-1][0])
	}
	return dp

	// deque
	//n := len(nums)
	//maxs := make([]int, n-k+1)
	//deque := make([]int, 0)
	//for i := 0; i < n; i++ {
	//	j := len(deque) - 1
	//	for j >= 0 && nums[i] > deque[j] {
	//		j--
	//	}
	//	deque = deque[:j+1]
	//	deque = append(deque, nums[i])
	//	if i >= k-1 {
	//		// 提前将下次窗体的出局的数移除，否则将难判断：nums[i-k+1]
	//		maxs[i-k+1] = deque[0]
	//		if deque[0] == nums[i-k+1] {
	//			deque = deque[1:]
	//		}
	//	}
	//}
	//return maxs

	// 暴力法
	//n := len(nums)
	//maxs := make([]int, n-k+1)
	//for i := 0; i < len(nums)-k+1; i++ {
	//	max := nums[i]
	//	for j := i + 1; j < i+k; j++ {
	//		max = MyTools.Max(nums[j], max)
	//	}
	//	maxs[i] = max
	//}
	//return maxs
}

//leetcode submit region end(Prohibit modification and deletion)
