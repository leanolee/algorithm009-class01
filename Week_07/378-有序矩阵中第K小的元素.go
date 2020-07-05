//给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。 
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。 
//
// 
//
// 示例： 
//
// matrix = [
//   [ 1,  5,  9],
//   [10, 11, 13],
//   [12, 13, 15]
//],
//k = 8,
//
//返回 13。
// 
//
// 
//
// 提示： 
//你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。 
// Related Topics 堆 二分查找
package main

import "fmt"

func main() {
	//matrix := [][]int{
	//	{1, 5, 9},
	//	{10, 11, 13},
	//	{12, 13, 15}}
	//k := 8
	matrix := [][]int{
		{1, 2},
		{1, 3}}
	k := 3
	smallest := kthSmallest(matrix, k)
	fmt.Println(smallest)

	//arr := []int{1, 3, 5, 8, 10, 13, 19}
	//num := 15
	//idx := insertIdx(arr, num)
	//fmt.Println(idx)
}

//leetcode submit region begin(Prohibit modification and deletion)
func kthSmallest(matrix [][]int, k int) int {
	// 二分：利用行、列有序
	n := len(matrix)
	//l, r := matrix[0][0], matrix[(k-1)/n][(k-1)%n]	// 范围缩小，行不通
	l, r := matrix[0][0], matrix[n-1][n-1]
	fmt.Println(r)
	for l < r {
		mid := (l + r) >> 1
		fmt.Println(mid, l, r)
		if check(matrix, k, n, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l

	// 归并：利用行有序

	// 堆

	// 全排序
}
func check(matrix [][]int, k int, n int, mid int) bool {
	for i := n - 1; i >= 0; i-- {
		l, r := 0, n
		for l < r {
			m := (l + r) >> 1
			if matrix[i][m] <= mid {
				l = m + 1
			} else {
				r = m
			}
		}
		k -= l
	}
	return k <= 0
}
func insertIdx(arr []int, val int) int {
	m, l, r := 0, 0, len(arr)
	for l < r {
		m = (l + r) >> 1
		if arr[m] <= val {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
