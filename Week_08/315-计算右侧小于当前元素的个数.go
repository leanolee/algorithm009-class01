//给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于 num
//s[i] 的元素的数量。 
//
// 示例: 
//
// 输入: [5,2,6,1]
//输出: [2,1,1,0] 
//解释:
//5 的右侧有 2 个更小的元素 (2 和 1).
//2 的右侧仅有 1 个更小的元素 (1).
//6 的右侧有 1 个更小的元素 (1).
//1 的右侧有 0 个更小的元素.
// 
// Related Topics 排序 树状数组 线段树 二分查找 分治算法
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 分治：归并
	nums := []int{5, 2, 6, 1}
	//nums := []int{-1, -1}
	smaller := countSmaller(nums)
	fmt.Println(smaller)
}

/*
同题型：lc-493
BIT参见：https://www.cnblogs.com/xenny/p/9739600.html
BIT：
	a：原数组排序后的数组
	c：树状数组
	ans：结果数组
	1.query：求 a[1 - i] 的和
	2.update：在 i 位置加上 k，此处 k=1
	3.遍历原数组，尾先入（后入的元素就可以判断，有多少个先入的元素是小于该后入元素）
		3.1.二分查找元素排序后的位置（首次出现）
		3.2.查询树状数组
		3.3.更新树状数组，为什么 idx+1？
			lowbit：取最后一位 1，所以需要 idx > 0
*/
//leetcode submit region begin(Prohibit modification and deletion)
func countSmaller(nums []int) []int {
	// BIT
	n := len(nums)
	ans, a, c := make([]int, n), make([]int, n), make([]int, n+1)
	copy(a, nums)
	sort.Ints(a)
	fmt.Println(a)
	for i := n - 1; i >= 0; i-- { // 3
		//idx := sort.SearchInts(a, nums[i]) // 3.1
		idx := binarySearch(a, 0, n, nums[i]) // 3.1
		ans[i] = query(c, idx)                // 3.2
		fmt.Println("idx:", idx, c)
		update(c, idx+1) // 3.3
		fmt.Println(c)
	}
	fmt.Println(c)
	return ans

	// 分治：归并
	//ans, kv := make([]int, len(nums)), make([][2]int, len(nums))
	//for i, v := range nums {
	//	kv[i][0], kv[i][1] = v, i
	//}
	//mergeSort(kv, 0, len(nums)-1, &ans)
	//return ans
}
func update(bit []int, i int) { // 2
	for i < len(bit) {
		bit[i]++
		i += i & -i
	}
}
func query(bit []int, i int) int { // 1
	sum := 0
	for i > 0 {
		sum += bit[i]
		i -= i & -i
	}
	return sum
}
func binarySearch(arr []int, l, r int, val int) int { // 3.1
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
func mergeSort(arr [][2]int, l int, r int, ans *[]int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(arr, l, mid, ans)
	mergeSort(arr, mid+1, r, ans)
	merge(arr, l, mid, r, ans)
}

func merge(arr [][2]int, l int, mid int, r int, ans *[]int) {
	temp, i, k := make([][2]int, r-l+1), l, 0
	for j := mid + 1; j <= r; j++ {
		for i <= mid && arr[i][0] <= arr[j][0] {
			temp[k], k, i = arr[i], k+1, i+1
		}
		temp[k], k = arr[j], k+1
		for c := i; c <= mid; c++ {
			(*ans)[arr[c][1]]++
		}
	}
	copy(arr[k+l:], arr[i:mid+1])
	copy(arr[l:], temp[:k])
}

//leetcode submit region end(Prohibit modification and deletion)
