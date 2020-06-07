//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 
//
// 示例: 
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
// Related Topics 回溯算法
package main

import "fmt"

func main() {
	n := 4
	k := 2
	i := combine(n, k)
	fmt.Println(i)
}

/*
子集：leetcode-78
	子集包含组合
*/
//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	//二进制排序
	combine := make([][]int, 0)
	nums := make([]int, k+1)
	i := 1
	for ; i <= k; i++ {
		nums[i-1] = i	// 初始化为：[1 2 0]
	}
	nums[i-1] = n + 1	// 设置上限：[1 2 5]
	for j := 0; j < k; {
		combine = append(combine, append([]int{}, nums[:k]...))
		j = 0
		for j < k && nums[j]+1 == nums[j+1] {	// 前一个数即将超过后一个
			nums[j] = j + 1	// 复位
			j++
		}
		nums[j]++
		/*
		[1 2 5]	// 初始值
		[1 3 5]
		[2 3 5]
		[1 4 5]
		[2 4 5]
		[3 4 5]
		[1 2 6]	// 新的一次循环
		*/
	}
	return combine

	// 子集：leetcode-78

	// 递归：回溯
	//combine := make([][]int, 0)
	//combineDFS(n, k, &combine, []int{}, 1)
	//return combine
}

func combineDFS(n int, k int, combine *[][]int, com []int, idx int) {
	plus := append([]int{}, com...)
	if len(com) == k {
		*combine = append(*combine, com)
		return
	}
	for i := idx; i <= n; i++ {
		plus = append(plus, i)
		combineDFS(n, k, combine, plus, i+1)
		plus = plus[:len(plus)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
