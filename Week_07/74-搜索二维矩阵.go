//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性： 
//
// 
// 每行中的整数从左到右按升序排列。 
// 每行的第一个整数大于前一行的最后一个整数。 
// 
//
// 示例 1: 
//
// 输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 13
//输出: false 
// Related Topics 数组 二分查找
package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 50
	b := searchMatrix(matrix, target)
	fmt.Println(b)
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	// 二分：O(log mn)
	m := len(matrix)
	if m == 0 {
		return false
	}
	l, r, n := 0, m*len(matrix[0])-1, len(matrix[0])
	for l <= r {
		mid := (l + r) >> 1
		val := matrix[mid/n][mid%n]
		switch {
		case val == target:
			return true
		case val < target:
			l = mid + 1
		case val > target:
			r = mid - 1
		}
	}
	return false

	// 层：O(m+n)
	if len(matrix) == 0 {
		return false
	}
	r, c := len(matrix)-1, len(matrix[0])-1
	for i, j := 0, c; i <= r && j >= 0; {
		val := matrix[i][j]
		switch {
		case val == target:
			return true
		case val < target:
			i, j = i+1, c
		case val > target:
			j--
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
