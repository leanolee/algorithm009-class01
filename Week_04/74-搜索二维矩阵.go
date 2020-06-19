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
	target := 23
	b := searchMatrix(matrix, target)
	fmt.Println(b)
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	// 二分
	//if len(matrix) == 0 {
	//	return false
	//}
	//r, c := len(matrix), len(matrix[0])
	//l, r := 0, r*c
	//for l < r {
	//	mid := (l + r) >> 1
	//	val := matrix[mid/c][mid%c]
	//	if val == target {
	//		return true
	//	} else if val < target {
	//		l = mid + 1
	//	} else { // val > target
	//		r = mid
	//	}
	//}
	//return false

	// 层
	if len(matrix) == 0 {
		return false
	}
	r, c := len(matrix), len(matrix[0])
	i, j := 0, c-1
	for i < r && j >= 0 {
		val := matrix[i][j]
		if val == target {
			return true
		} else if val < target {
			i++
			j = c - 1
		} else {
			j--
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
