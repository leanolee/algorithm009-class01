//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。 
//
// 
//
// 示例 1: 
//
// 输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
// 
//
// 示例 2: 
//
// 输入: nums = [1], k = 1
//输出: [1] 
//
// 
//
// 提示： 
//
// 
// 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。 
// 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。 
// 你可以按任意顺序返回答案。 
// 
// Related Topics 堆 哈希表
package main

import "fmt"

func main() {
	nums := []int{0, 6, 1, 1, 1, 2, 2, 3, 4, 4, 4, 1, 3, 3, 3}
	k := 3
	frequent := topKFrequent(nums, k)
	fmt.Println(frequent)
}

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	// 小顶堆
	heap, memo := make([]int, k), make(map[int]int)
	for _, v := range nums {
		memo[v]++
	}
	i := 0
	for key, val := range memo {
		if i == k {
			if val > memo[heap[0]] {
				del(heap, k, memo, key, val)
			}
		} else {
			insert(heap, i, memo, key, val)
			i++
		}
	}
	return heap
}

func del(heap []int, k int, memo map[int]int, key int, val int) {
	i := 0
	for down := i<<1 + 1; down < k; i, down = down, down<<1+1 {
		if down+1 < k && memo[heap[down+1]] < memo[heap[down]] {
			down++
		}
		if val < memo[heap[down]] {
			break
		}
		heap[i] = heap[down]
	}
	heap[i] = key
}
func insert(heap []int, i int, memo map[int]int, key, val int) {
	for up := (i - 1) >> 1; up >= 0 && memo[heap[up]] > val; i, up = up, (up-1)>>1 {
		heap[i] = heap[up]
	}
	heap[i] = key
}

//leetcode submit region end(Prohibit modification and deletion)
