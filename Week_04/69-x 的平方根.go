//实现 int sqrt(int x) 函数。 
//
// 计算并返回 x 的平方根，其中 x 是非负整数。 
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。 
//
// 示例 1: 
//
// 输入: 4
//输出: 2
// 
//
// 示例 2: 
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842..., 
//     由于返回类型是整数，小数部分将被舍去。
// 
// Related Topics 数学 二分查找
package main

import "fmt"

func main() {
	x := 8
	sqrt := mySqrt(x)
	fmt.Println(sqrt)
}

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	// 牛顿迭代法
	//g := x
	//for g*g > x {
	//	g = (g + x/g) >> 1
	//}
	//return g

	// 二分
	l, r := 1, x
	res := 0
	for l <= r {
		mid := (l + r) >> 1
		if mid*mid <= x {
			res, l = mid, mid+1
		} else {
			r = mid - 1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
