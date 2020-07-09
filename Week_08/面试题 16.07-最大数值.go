//编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。 
// 示例： 
// 输入： a = 1, b = 2
//输出： 2
// 
// Related Topics 位运算 数学
package main

import (
	"fmt"
	"math"
)

func main() {
	//a, b := 13, 9
	//a, b := 16, 4
	//a, b := 3, 8
	a, b := 2147483647, -2147483648
	i := maximum(a, b)
	fmt.Println(i)
	fmt.Println(math.MinInt32 & math.MaxInt32)
}

/*
实质：(a+b+abs(a-b))>>1
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maximum(a int, b int) int {
	// 位运算：64
	x, y := int64(a), int64(b)
	sum, diff := x+y, x-y
	diff = diff ^ (diff >> 63) - diff>>63 // 如果diff是负数：diff >> 63 = -1
	return int((sum + diff) >> 1)

	//c := a ^ b // 获得最高位的 1
	//c = c | (c >> 1)
	//c = c | (c >> 2)
	//c = c | (c >> 4)
	//c = c | (c >> 8)
	//c = c | (c >> 16)
	//c = c>>1 + 1
	///*
	//	假设 a > b：
	//		c&a = c
	//		c&b = 0
	//		^(c&a-1)&a：取大于等于 c 的 a 的位
	//		^(c&a-1)&a ^ a：取小于 c 的 a 的位，不对！！！
	//	痛点：
	//		1.取小于 c 的 a 的位
	//		2.当存在负数情况：如 a, b := 2147483647, -2147483648
	//*/
	//return ^(c&a-1)&a | ((c&a-1)&a ^ a) | ^(c&b-1)&b | ((c&b-1)&b ^ b)
}

//leetcode submit region end(Prohibit modification and deletion)
