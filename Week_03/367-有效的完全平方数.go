//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。 
//
// 说明：不要使用任何内置的库函数，如 sqrt。 
//
// 示例 1： 
//
// 输入：16
//输出：True 
//
// 示例 2： 
//
// 输入：14
//输出：False
// 
// Related Topics 数学 二分查找
package main

import "fmt"

func main() {
	num := 81
	square := isPerfectSquare(num)
	fmt.Println(square)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	// 牛顿迭代法

	// 二分：迭代
	var mid, square int
	for i, j := 1, num; i <= j; {
		mid = (i + j) >> 1
		square = mid * mid
		switch {
		case square > num:
			j = mid - 1
		case square < num:
			i = mid + 1
		case square == num:
			return true
		}
	}
	return false

	// 二分
	//return isPerfectSquareRecursion(num, 1, num)
}

func isPerfectSquareRecursion(num int, i int, j int) bool {
	if i > j {
		return false
	}
	mid := (i + j) >> 1
	square := mid * mid
	switch {
	case square > num:
		return isPerfectSquareRecursion(num, i, mid-1)
	case square < num:
		return isPerfectSquareRecursion(num, mid+1, j)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
