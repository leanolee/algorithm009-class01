//不使用运算符 + 和 - ，计算两整数 a 、b 之和。 
//
// 示例 1: 
//
// 输入: a = 1, b = 2
//输出: 3
// 
//
// 示例 2: 
//
// 输入: a = -2, b = 3
//输出: 1 
// Related Topics 位运算
package main

import "fmt"

func main() {
	a, b := 3, -2
	sum := getSum(a, b)
	fmt.Println(sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func getSum(a int, b int) int {
	// 位运算
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
