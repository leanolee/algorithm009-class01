//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。 
//
// 
//
// 示例 1： 
//
// 输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
// 
//
// 示例 2： 
//
// 输入：arr = [0,1,2,1], k = 1
//输出：[0] 
//
// 
//
// 限制： 
//
// 
// 0 <= k <= arr.length <= 10000 
// 0 <= arr[i] <= 10000 
// 
// Related Topics 堆 分治算法
package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 1}
	k := 2
	numbers := getLeastNumbers(arr, k)
	fmt.Println(numbers)
}

/*
暴力法：O(kn),O(1)
二叉堆（小顶）：O(nlogn),O(n)

API排序：O(nlogn),O(1)
TOP-K：
*/
//leetcode submit region begin(Prohibit modification and deletion)
func getLeastNumbers(arr []int, k int) []int {
	// TOP-K：快排

	// API排序
	//sort.Ints(arr)
	//return arr[:k]

	// 二叉堆：小顶，元素个数为k

	// 二叉堆：小顶
	n := len(arr)
	numbers := make([]int, k)
	binaryHeap := make([]int, n)
	for i, a := range arr {
		// insert
		idx := i
		for idx > 0 && a < binaryHeap[(idx-1)>>1] {
			binaryHeap[idx] = binaryHeap[(idx-1)>>1]
			idx = (idx - 1) >> 1
		}
		binaryHeap[idx] = a
	}
	for i, _ := range numbers {
		numbers[i] = binaryHeap[0]
		// delete heap top
		idx := 0
		n := len(binaryHeap)
		tail := binaryHeap[n-1]	// 尾节点提到根
		for {
			nextIdx := idx<<1 + 1 // 默认左
			if nextIdx >= n {     // 左右都没
				break
			}
			if nextIdx+1 < n && binaryHeap[nextIdx+1] < binaryHeap[nextIdx] { // nextIdx+1 == n：只有左
				nextIdx++
			}
			if tail < binaryHeap[nextIdx] { // 已经满足（小顶）二叉堆
				break
			}
			binaryHeap[idx] = binaryHeap[nextIdx]
			idx = nextIdx
		}
		binaryHeap[idx] = tail
		binaryHeap = binaryHeap[:n-1]
	}
	return numbers

	// 暴力法：优化
	//numbers := make([]int, k)
	//for i := 0; i < k; i++ {
	//	minIdx := i
	//	for j := i + 1; j < len(arr); j++ {
	//		if arr[j] < arr[minIdx] {
	//			minIdx = j
	//		}
	//	}
	//	numbers[i] = arr[minIdx]
	//	arr[i], arr[minIdx] = arr[minIdx], arr[i]
	//}
	//return numbers

	// 暴力法：超时
	//numbers := make([]int, k)
	//idxs := make([]int, k)
	//for i := 0; i < k; i++ {
	//	minIdx := -1
	//out:
	//	for j := 0; j < len(arr); j++ {
	//		for k := 0; k < i; k++ {
	//			if idxs[k] == j {
	//				continue out
	//			}
	//		}
	//		if minIdx == -1 || arr[j] < arr[minIdx] {
	//			minIdx = j
	//		}
	//	}
	//	idxs[i] = minIdx
	//	numbers[i] = arr[minIdx]
	//}
	//return numbers
}

//leetcode submit region end(Prohibit modification and deletion)
