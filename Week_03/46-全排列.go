//给定一个 没有重复 数字的序列，返回其所有可能的全排列。 
//
// 示例: 
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//] 
// Related Topics 回溯算法
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	i := permute(nums)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	// 递归：回溯
	permute := make([][]int, 0)
	permuteDFS(nums, &permute, 0)
	return permute

	// 递归：切数组，很慢
	//permute := make([][]int, 0)
	//permuteDFS_(nums, &permute, []int{})
	//return permute
}

func permuteDFS(nums []int, permute *[][]int, idx int) {
	if idx == len(nums)-1 { // 其实 idx==len(nums)-1 就行了，因为阶乘最后一次乘 1
		*permute = append(*permute, append([]int{}, nums...))
		return
	}
	/*
		阶乘：3*2*1
		idx=0:3个分支
			第一次交换，固定第一个元素：* 3
		idx=1:6个分支
			第二次交换，固定第二个元素，在第一次的基础上再分支：* 2
		idx=2:6个分支
			第3个元素已经固定了，所以可以省略这步：* 1
	*/
	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		permuteDFS(nums, permute, idx+1)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}

func permuteDFS_(nums []int, permute *[][]int, ints []int) {
	if len(nums) == 0 {
		*permute = append(*permute, ints)
		return
	}
	for i := 0; i < len(nums); i++ {
		//numsNext := append([]int{}, nums[:i]...)
		//permuteDFS(append(numsNext, nums[i+1:]...), permute, append(ints, nums[i]))
		permuteDFS_(append(append([]int{}, nums[:i]...), nums[i+1:]...), permute, append(ints, nums[i]))
	}
}

//leetcode submit region end(Prohibit modification and deletion)
