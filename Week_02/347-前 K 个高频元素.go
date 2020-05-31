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

import (
	"fmt"
)

func main() {
	//nums := []int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}
	nums := []int{0, 0, 0, 0, 2, 2, 2, 3, 3, 6, 13, 23, 32, 34, 44, 45, 54, 76, 654, 765, 889, 987}
	//nums := []int{1}
	k := 3
	frequent := topKFrequent(nums, k)
	//fmt.Println(len(nums)) // 22->16, count 14
	fmt.Println(frequent)

	j := 0
	//i := (j - 1) / 2
	i := (j - 1) >> 1
	fmt.Println(i)
}

/*
暴力法：O(nk),O(n)
	1.map 统计每个数的次数
	2.每次遍历取出top 1
API 1：O(nlogn),O(n)
API 2：O(nlogn),O(n)
大顶堆：
小顶堆：O(nlogk),O(n)
快排：
*/
//leetcode submit region begin(Prohibit modification and deletion)
type TopK struct {
	value int
	times int
}
type topKs []TopK

func (t topKs) Len() int {
	return len(t)
}
func (t topKs) Less(i int, j int) bool {
	return t[i].times > t[j].times
}
func (t topKs) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}
func deleteAndInsert(k int, topK TopK, heap []TopK) {
	// go 源码
	heap[0] = topK
	index := 0
	for {
		downIdx := 2*index + 1
		if downIdx >= k {
			break
		}
		if downIdx < k-1 && heap[downIdx+1].times < heap[downIdx].times {
			downIdx++
		}
		if heap[index].times < heap[downIdx].times {
			break
		}
		heap[downIdx], heap[index] = heap[index], heap[downIdx]
		index = downIdx
	}

	// 个人写法
	//index, downIdx := 0, 0
	//// 三种情况：1.左、右节点都不存在，2.只左节点存在，3.左、右节点都存在
	//for {
	//	downIdx = index<<1 + 1 // 左节点
	//	if downIdx >= k {      // 1
	//		break
	//	}
	//	if downIdx < k-1 && heap[downIdx+1].times < heap[downIdx].times { // 3，2：downIdx == k-1 只有左
	//		downIdx++ // 右节点小，替换右节点
	//	}
	//	if topK.times < heap[downIdx].times { // 已经满足（小顶）二叉堆
	//		break
	//	}
	//	heap[index] = heap[downIdx]
	//	index = downIdx
	//}
	//heap[index] = topK
}
func insert(index int, topK TopK, heap []TopK) {
	// go 源码：但是速度没自己写的快
	heap[index] = topK
	for {
		upIdx := (index - 1) / 2
		// 注意这里 upIdx == index
		if upIdx == index || heap[index].times > heap[upIdx].times {
			break
		}
		heap[upIdx], heap[index] = heap[index], heap[upIdx]
		index = upIdx
	}

	// 个人写法
	//upIdx := (index - 1) >> 1
	//for ; upIdx >= 0 && topK.times < heap[upIdx].times; index, upIdx = upIdx, (upIdx-1)>>1 {
	//	heap[index] = heap[upIdx]
	//}
	//heap[index] = topK
}
func qSort(arr []TopK, begin int, end int, k int) int {
	// 有返回值
	if begin >= end {
		return 0
	}
	pivot := qPartition(arr, begin, end)
	pivotIdx := 0
	switch count := end - pivot + 1; {
	case count == k:
		pivotIdx = pivot
	case count == k+1:
		pivotIdx = pivot + 1
	case count > k+1:
		pivotIdx = qSort(arr, pivot+1, end, k)
	case count < k:
		pivotIdx = qSort(arr, begin, pivot-1, k-count)
	}
	return pivotIdx

	// 无返回值
	//if begin >= end {
	//	return
	//}
	//pivot := qPartition(arr, begin, end)
	//count := end - pivot + 1
	//if count == k {
	//	pivotIdx = pivot
	//	return
	//} else if count == k+1 {
	//	pivotIdx = pivot + 1
	//	return
	//} else if count > k+1 {
	//	qSort(arr, pivot+1, end, k)
	//} else {
	//	qSort(arr, begin, pivot-1, k-count)
	//}
}

//var pivotIdx = 0

func qPartition(arr []TopK, begin int, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if arr[i].times < arr[pivot].times {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}
	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
}
func topKFrequent(nums []int, k int) []int {
	// 快排
	frequent := make([]int, k)
	hash := make(map[int]int) // k:num,v:times
	for _, num := range nums {
		hash[num]++
	}
	quick := make([]TopK, len(hash))
	i := 0
	for key, val := range hash {
		quick[i] = TopK{key, val}
		i++
	}
	pivotIdx := qSort(quick, 0, i-1, k)	// 无返回值的情况，有个坑：leetcode重复测试时，会沿用上次测试后的全局变量的值
	//fmt.Println("pivotIdx:", pivotIdx)
	for j := pivotIdx; j < i; j++ {
		frequent[j-pivotIdx] = quick[j].value
	}
	return frequent

	// 小顶堆
	//frequent := make([]int, k)
	//hash := make(map[int]int) // k:num,v:times
	//for _, num := range nums {
	//	hash[num]++
	//}
	//minHeap := make([]TopK, k) // 存 k,比较 v
	//count := 0
	//for key, value := range hash {
	//	if count < k {
	//		insert(count, TopK{key, value}, minHeap) // insert
	//		count++
	//		fmt.Println(minHeap)
	//	} else { // count == k
	//		if value > minHeap[0].times {
	//			deleteAndInsert(count, TopK{key, value}, minHeap) // delete and insert
	//		}
	//	}
	//}
	//for i, topK := range minHeap {
	//	frequent[i] = topK.value
	//}
	//return frequent

	// 大顶堆

	// Sort by API 2：速度最快的
	//frequent := make([]int, k)
	//hash := make(map[int]int)
	//for _, num := range nums {
	//	hash[num]++
	//}
	//topK := topKs(make([]TopK, len(hash)))
	//index := 0
	//for k, v := range hash {
	//	topK[index] = TopK{k, v}
	//	index++
	//}
	//sort.Sort(topK)
	//fmt.Println(topK)
	//for i, _ := range frequent {
	//	frequent[i] = topK[i].value
	//}
	//return frequent

	// Sort by API 1
	//sort.Ints(nums)
	//frequent := make([]int, k)
	//topK := topKs{{nums[0], 1}}
	//index := 0
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] == nums[i-1] {
	//		topK[index].times++
	//	} else {
	//		topK = append(topK, TopK{nums[i], 1})
	//		index++
	//	}
	//}
	//sort.Sort(topK)
	//for i, _ := range frequent {
	//	frequent[i] = topK[i].value
	//}
	//return frequent

	// 暴力法
	//frequent := make([]int, k)
	//hash := make(map[int]int)
	//for _, num := range nums {
	//	hash[num]++
	//}
	//for i, _ := range frequent {
	//	max := 0
	//	num := 0
	//	for k, v := range hash {
	//		if max < v {
	//			max = v
	//			num = k
	//		}
	//	}
	//	delete(hash, num)
	//	frequent[i] = num
	//}
	//return frequent
}

//leetcode submit region end(Prohibit modification and deletion)
