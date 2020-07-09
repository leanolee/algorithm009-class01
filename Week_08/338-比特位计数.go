//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。 
//
// 示例 1: 
//
// 输入: 2
//输出: [0,1,1] 
//
// 示例 2: 
//
// 输入: 5
//输出: [0,1,1,2,1,2] 
//
// 进阶: 
//
// 
// 给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？ 
// 要求算法的空间复杂度为O(n)。 
// 你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。 
// 
// Related Topics 位运算 动态规划
package main

import "fmt"

func main() {
	num := 100
	bits := countBits(num)
	fmt.Println(bits)
}

/*
dp：
	前一个数是偶数：dp[i] = dp[i-1]+1
	前一个数是奇数：dp[i] = dp[i-1] - k + 1
		前一个数尾有 k 个连续的 1
*/
//leetcode submit region begin(Prohibit modification and deletion)
func countBits(num int) []int {
	// 最佳二
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp

	// 最佳一
	//dp := make([]int, num+1)
	//for i := 1; i <= num; i++ {
	//	dp[i] = dp[i&(i-1)] + 1
	//}
	//return dp

	// dp：个人
	//dp := make([]int, num+1)
	//for i := 1; i <= num; i++ {
	//	dp[i] = dp[i-1] + 1
	//	if i&1 == 0 { // 当前为奇数，前一个是偶数
	//		count := 0
	//		for j := i - 1; j&1 == 1; j >>= 1 {	// 此处可以用二分吧
	//			count++
	//		}
	//		dp[i] -= count
	//	}
	//}
	//return dp
}

//leetcode submit region end(Prohibit modification and deletion)
