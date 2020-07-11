//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。 
//
// 说明： 
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？ 
//
// 示例 1: 
//
// 输入: [2,2,3,2]
//输出: 3
// 
//
// 示例 2: 
//
// 输入: [0,1,0,1,0,1,99]
//输出: 99 
// Related Topics 位运算
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	number2 := singleNumber2(nums)
	fmt.Println(number2)
}

/*
异或：https://leetcode-cn.com/problems/single-number-ii/solution/zhi-chu-xian-yi-ci-de-shu-zi-ii-by-leetcode/
	1.bit 出现 1 次数：
		奇数：1
		偶数：0
	2.位掩码 seen_once 仅保留出现一次的数字，不保留出现三次的数字
		seen_once = ~seen_twice & (seen_once ^ num)
		seen_twice = ~seen_once & (seen_twice ^ num)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber2(nums []int) int {
	// 位运算
	once, twice := 0, 0
	for _, n := range nums {
		once = ^twice & (once ^ n)
		twice = ^once & (twice ^ n)
	}
	return once

	// hash：3×(a+b+c)−(a+a+a+b+b+b+c)=2c
	//memo := make(map[int]int)
	//for _, n := range nums {
	//	if memo[n] == 2 {
	//		delete(memo, n)
	//	} else {
	//		memo[n]++
	//	}
	//}
	//for k, _ := range memo {
	//	return k
	//}
	//return 0
}

//leetcode submit region end(Prohibit modification and deletion)
