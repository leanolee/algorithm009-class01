//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。 
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//] 
// Related Topics 位运算 数组 回溯算法
package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 2, 3, 4}
	i := subsets(nums)
	fmt.Println(i)
}

/*
DFS：迭代
DFS：回溯
位运算：
DP：递归
DP：迭代
*/
//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	// dp：
	//subs := [][]int{{}}
	//for i := 0; i < len(nums); i++ {
	//	n := len(subs)
	//	for j := 0; j < n; j++ {
	//		plus := append(make([]int, 0), subs[j]...)
	//		subs = append(subs, append(plus, nums[i]))
	//	}
	//}
	//return subs

	// 位运算
	//n := len(nums)
	//count := 1 << n
	//subs := make([][]int, count)
	//for i := 0; i < count; i++ {
	//	plus := make([]int, 0)
	//	//for j, idx := i, n-1; j > 0 && idx >= 0; j, idx = j>>1, idx-1 {
	//	for j, idx := i, n-1; j > 0; j, idx = j>>1, idx-1 {
	//		if j&1 == 1 {
	//			plus = append(plus, nums[idx])
	//		}
	//	}
	//	subs[i] = plus
	//}
	//return subs

	// 递归：思路同dp
	//subs := [][]int{{}}
	//subsetsDP(0, &subs, nums)
	//return subs

	// DFS：回溯 写法二
	subs := [][]int{{}}
	subsetsDFS(0, []int{}, &subs, nums)
	return subs

	//DFS：回溯 写法一
	//subs := make([][]int, 0)
	//subsetsDFS(0, []int{}, &subs, nums)
	//return subs

	// dfs：2
	//subs := [][]int{{}}
	//subsetsDFS(0, []int{}, &subs, nums)
	//return subs

	// dfs：1
	//subs := [][]int{{}}
	//subsetsRecursion([]int{}, &subs, nums)
	//return subs
}

func subsetsDP(idx int, subs *[][]int, nums []int) {
	if idx == len(nums) {
		return
	}
	n := len(*subs)
	for i := 0; i < n; i++ {
		plus := append([]int{}, (*subs)[i]...)
		*subs = append(*subs, append(plus, nums[idx]))
	}
	subsetsDP(idx+1, subs, nums)
}
func subsetsDFS(idx int, sub []int, subs *[][]int, nums []int) {
	// DFS 回溯：写法二
	if idx == len(nums) {
		return
	}
	subsetsDFS(idx+1, sub, subs, nums) // 不选择
	plus := append(make([]int, 0), sub...)
	plus = append(plus, nums[idx])
	*subs = append(*subs, plus)
	subsetsDFS(idx+1, plus, subs, nums) // 选择
	plus = plus[:len(plus)-1]

	// DFS 回溯：写法一
	//plus := make([]int, len(sub))	// 有公共数据，回溯
	//copy(plus, sub)
	//*subs = append(*subs, plus)
	//for ; idx < len(nums); idx++ {
	//	plus = append(plus, nums[idx])
	//	subsetsDFS(idx+1, plus, subs, nums)
	//	plus = plus[:len(plus)-1]
	//}

	// dfs 写法二 通过传递 idx
	//if idx == len(nums) {
	//	return
	//}
	//for i := idx; i < len(nums); i++ {
	//	plus := append([]int{}, sub...)
	//	plus = append(plus, nums[i])
	//	*subs = append(*subs, plus)
	//	subsetsDFS(i+1, plus, subs, nums)
	//	//plus = plus[:len(plus)-1]	// 没有改变共有数据，不需要回溯
	//}
}
func subsetsRecursion(ints []int, subs *[][]int, nums []int) {
	// dfs 写法一
	if len(nums) == 0 {
		return
	}
	for i := 0; i < len(nums); i++ {
		plus := append([]int{}, ints...) // 一定要新建，不然提交错误。应该是判断了指针值
		plus = append(plus, nums[i])
		*subs = append(*subs, plus)
		subsetsRecursion(plus, subs, nums[i+1:])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
