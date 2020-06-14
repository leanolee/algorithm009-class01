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
	num := 14
	square := isPerfectSquare(num)
	fmt.Println(square)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	// 位运算
	//return num != 7 && num&-num == 0
	//return num != 7 && num&(num-1) == 0

	// dp
	if num < 2 {
		return true
	}
	x := num >> 1
	g := 1
	for i := 1; i <= x; i++ {
		g += i<<1 + 1
		if g < num {
			continue
		}
		return g == num
	}
	return false

	// 二分
	//l, r := 1, num
	//for l <= r {
	//	mid := (l + r) >> 1
	//	pow := mid * mid
	//	if pow == num {
	//		return true
	//	} else if pow < num {
	//		l = mid + 1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return false
}

//leetcode submit region end(Prohibit modification and deletion)
