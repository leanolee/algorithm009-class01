//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。 
//
// 说明： 
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？ 
//
// 示例 1: 
//
// 输入: [2,2,1]
//输出: 1
// 
//
// 示例 2: 
//
// 输入: [4,1,2,1,2]
//输出: 4 
// Related Topics 位运算 哈希表
package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	number0 := singleNumber1(nums)
	fmt.Println(number0)
}

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber1(nums []int) int {
	// 位运算
	//ans := 0
	//for _, n := range nums {
	//	ans ^= n
	//}
	//return ans

	// hash
	memo := make(map[int]bool)
	for _, n := range nums {
		if memo[n] {
			delete(memo, n)
		} else {
			memo[n] = true
		}
	}
	for k, _ := range memo {
		return k
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
