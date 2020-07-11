//给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。 
//
// 你需要返回给定数组中的重要翻转对的数量。 
//
// 示例 1: 
//
// 
//输入: [1,3,2,3,1]
//输出: 2
// 
//
// 示例 2: 
//
// 
//输入: [2,4,3,5,1]
//输出: 3
// 
//
// 注意: 
//
// 
// 给定数组的长度不会超过50000。 
// 输入数组中的所有数字都在32位整数的表示范围内。 
// 
// Related Topics 排序 树状数组 线段树 二分查找 分治算法
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2, 3, 1}
	//nums := []int{2, 4, 3, 5, 1}
	pairs := reversePairs(nums)
	fmt.Println(pairs, nums)

	//sort.Ints(nums)
	//fmt.Println(nums)
	//search := binarySearch(nums, 0, len(nums), 2)
	//fmt.Println(search)
}

/*
同题型：lc-315
归并排序：
	1.分治
		count：统计符合题意 arr[idx] > 2*arr[j] 的个数
		i：用于合并数组时，左侧数组的指针
		j：用于合并数组时，右侧数组的指针；且用于查找符合 arr[idx] > 2*arr[j] 时，右侧数组的指针
		k：合并后的数组的指针
		idx：用于查找符合 arr[idx] > 2*arr[j] 时，左侧数组的指针
	2.归并，合并两个有序数组：左侧数组 l - mid；右侧数组 mid+1 - r
		2.1.合并左侧数组
		2.2.合并右侧数组
		2.3.将左侧数组剩余的未合并的“大数”移至 temp 的末尾（由于temp最终要考到到原数组中，所以可以直接移至原数组 l+k: 的位置）
	3.查询左侧数组中符合 arr[idx] > 2*arr[j] 的数
		3.1.找出当前左侧数组中，符合 arr[idx] > 2*arr[j] 的idx位置；arr[idx+1],arr[idx+2],...,arr[mid]，都符合 > 2*arr[j]
			arr[idx]+1)>>1 <= arr[j]：+1 降低溢出风险（偷个懒）
		3.2.累加 arr[idx] > 2*arr[j] 个数：idx-mid 之间的数，都符合 > 2*arr[j]
	4.将合并后的数组拷贝回原数组
BIT：https://leetcode.com/problems/reverse-pairs/discuss/97272/Clean-Java-Solution-using-Enhanced-Binary-Search-Tree
BST：
	1.构建struct并创建root，默认count=1
	2.查询BST
		2.1.到叶子节点
		2.2.节点 val <= 目标，说明要往 left 节点查询
		2.3.节点 val > 目标，累加左子树的个数，并往 right 节点查询
	3.BST插入元素
		3.1.要插入的节点不存在，新建节点并返回该节点
		3.2.存在，返回该节点
		3.3.插入 val > 节点val，往右插入
		3.4.插入 val < 节点val，节点count++（左子树节点个数+1），往左插入
*/
//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	// 树状数组：BIT：https://leetcode.com/problems/reverse-pairs/discuss/162757/Python-BIT-using-ranks-Clear-O(nlog(n))
	cache := make(map[int]struct{})
	for _, v := range nums {
		cache[v] = struct{}{}
		cache[v<<1] = struct{}{}
	}
	n := len(cache)
	a, c, count, i := make([]int, n), make([]int, n+1), 0, 0
	for k, _ := range cache {
		a[i], i = k, i+1
	}
	sort.Ints(a)
	//fmt.Println(a)
	for i := len(nums) - 1; i >= 0; i-- { // 重点：为什么用 idx1, idx2 区分？还没完全想明白
		idx1, idx2 := binarySearch(a, 0, n, nums[i]), binarySearch(a, 0, n, nums[i]<<1)
		count += query(c, idx1)
		//fmt.Println(i, idx1, c)
		update(c, idx2+1)
	}
	return count

	// 归并排序
	//return mergeSort_(nums, 0, len(nums)-1)

	// BST：超时，因为数组如果单调递增，就退化为链表
	// 参考：https://leetcode.com/problems/reverse-pairs/discuss/97272/Clean-Java-Solution-using-Enhanced-Binary-Search-Tree
	//ans, n := 0, len(nums)
	//if n > 0 {
	//	root := &Node{nil, nil, nums[n-1], 1} // 1
	//	for i := n - 2; i >= 0; i-- {
	//		ans += queryBST(root, (nums[i]+1)>>1) // 2
	//		insertBST(root, nums[i])              // 3
	//	}
	//}
	//return ans
}
func binarySearch(arr []int, l, r int, val int) int {
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
func update(bit []int, i int) {
	for i < len(bit) {
		bit[i] ++
		i += i & -i
	}
}
func query(bit []int, i int) int {
	count := 0
	for i > 0 {
		count += bit[i]
		i -= i & -i
	}
	return count
}
func mergeSort_(arr []int, l int, r int) int {
	if l >= r {
		return 0
	}
	count, mid := 0, (l+r)>>1
	count = mergeSort_(arr, l, mid) + mergeSort_(arr, mid+1, r) // 1
	temp, i, k := make([]int, r-l+1), l, 0
	for j, idx := mid+1, l; j <= r; j++ { // 2
		for ; i <= mid && arr[i] < arr[j]; i++ {
			temp[k], k = arr[i], k+1 // 2.1
		}
		temp[k], k = arr[j], k+1                      // 2.2
		for idx <= mid && (arr[idx]+1)>>1 <= arr[j] { // 3
			idx++ // 3.1
		}
		count += mid - idx + 1 // 3.2
	}
	copy(arr[l+k:], arr[i:mid+1]) // 2.3
	copy(arr[l:], temp[:k])       // 4
	return count

	// 错误！
	//temp, j, k := make([]int, r-l+1), mid+1, 0
	//for i, idx := l, l; i <= mid; i++ {
	//	for idx <= mid && (arr[idx]+1)>>1 <= arr[j] {
	//		idx++
	//	}
	//	count += mid - idx + 1 // 会导致 i 循环时，重复计算：统计左侧，则右侧元素每次都要变；统计右侧，则左侧元素每次都要变
	//	//fmt.Println(count)
	//	for ; j <= r && arr[i] > arr[j]; j++ {
	//		temp[k], k = arr[j], k+1
	//	}
	//	temp[k], k = arr[i], k+1
	//}
	//copy(arr[l:], temp[:k])
	//return count
}
func insertBST(root *Node, val int) *Node {
	if root == nil { // 3.1
		return &Node{nil, nil, val, 1}
	}
	switch {
	case root.val == val: // 3.2
		root.count++
	case root.val < val: // 3.3
		root.r = insertBST(root.r, val)
	case root.val > val: // 3.4
		root.count++
		root.l = insertBST(root.l, val)
	}
	return root
}
func queryBST(root *Node, half int) int {
	if root == nil { // 2.1
		return 0
	}
	if root.val >= half { // 2.2
		return queryBST(root.l, half)
	} else { // 2.3
		return root.count + queryBST(root.r, half)
	}
}

type Node struct {
	l, r       *Node
	val, count int
}

//leetcode submit region end(Prohibit modification and deletion)
