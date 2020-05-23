//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。 
//
// 示例: 
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0] 
//
// 说明: 
//
// 
// 必须在原数组上操作，不能拷贝额外的数组。 
// 尽量减少操作次数。 
// 
// Related Topics 数组 双指针
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

/*
补全法：O(n),O(1)
	思路：遇到非0数据，则往前 copy，最后补全0
双指针：O(n),O(1)
	思路：循环的指针固定当前遍历的位置（遇到非0则copy），另外的指针指向下一个 copy 的位置
*/
//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	//双指针
	//for i, j := 0, 0; i < len(nums); i++ {
	//	if nums[i] != 0 {
	//		nums[i], nums[j] = nums[j], nums[i]
	//		j++
	//	}
	//}

	// 双指针：但是这样变了顺序，不符合要求
	//for i, j := 0, len(nums)-1; i < j; {
	//	for nums[j] == 0 && j > i {
	//		j--
	//	}
	//	for nums[i] != 0 && j > i {
	//		i++
	//	}
	//	nums[i], nums[j] = nums[j], nums[i]
	//	i++
	//	j--
	//}

	// 补全法
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; i > j; i-- {
		nums[i-1] = 0
	}

	// 开辟数组
	//arr := make([]int, len(nums))
	//j := 0
	//for i, v := range nums {
	//	if v != 0 {
	//		arr[i] = v
	//		j++
	//	}
	//}
	//nums = arr
}

//leetcode submit region end(Prohibit modification and deletion)
