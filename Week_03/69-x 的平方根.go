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

import (
	"fmt"
	"math"
)

func main() {
	x := 16
	sqrt := mySqrt(x)
	fmt.Println(sqrt)
}

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	// 二分法

	return int(InvSqrt(float32(x)))	// 太精准，提交不成功
}
func InvSqrt(x float32) float32 {
	var xhalf float32 = 0.5 * x // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	return 1 / x
}

/*
16 -> 4.0000005
16 -> 4.006779
*/
func invSqrt(x float32) float32 { // 使用 float64 会运算成 0
	xhalf := 0.5 * x            // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float	(or i = 0x5f3759df - (i >> 1))
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)   // 运行3次，会非常精确
	return 1 / x
}

//leetcode submit region end(Prohibit modification and deletion)
