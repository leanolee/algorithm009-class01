//给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。 
//
// 示例: 
//
// 输入: 38
//输出: 2 
//解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
// 
//
// 进阶: 
//你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？ 
// Related Topics 数学
package main

import "fmt"

func main() {
	num := 3923
	digits := addDigits(num)
	fmt.Println(digits)
}
/*
参见：https://leetcode-cn.com/problems/add-digits/solution/san-bu-qing-song-li-jie-o1-by-data-t/
方法一：除模
	时间复杂度：O(n)，n为数字的位数
方法二：模9，方法一的进阶
	1. 假设 a < 10，分析 a % 9 = ? 如何得到结果仍然是 a
		即 a % 9 = a
		(a-1)%9+1 = a
	2. 分析任意数：num
		= a1 + a2*10    + a3*100   + ... + an*(10*n)
		= a1 + a2*(9+1) + a3(99+1) + ... + an*(10*n-1+1)
		每项都 %9：
		= a1 + a2 + a3 + ... + an（正符合该题，如果 num > 10，重复上面操作）
	3. a1 ~ an 都是 <10 的数，结合第1步的结论
		return (num-1)%9 + 1
	时间复杂度：O(1)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func addDigits(num int) int {
	return (num-1)%9 + 1

	//for num >= 10 {
	//	num = num%10 + num/10
	//}
	//return num
}

//leetcode submit region end(Prohibit modification and deletion)
