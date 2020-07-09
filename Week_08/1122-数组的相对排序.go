//给你两个数组，arr1 和 arr2， 
//
// 
// arr2 中的元素各不相同 
// arr2 中的每个元素都出现在 arr1 中 
// 
//
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末
//尾。 
//
// 
//
// 示例： 
//
// 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//输出：[2,2,2,1,4,3,3,9,6,7,19]
// 
//
// 
//
// 提示： 
//
// 
// arr1.length, arr2.length <= 1000 
// 0 <= arr1[i], arr2[i] <= 1000 
// arr2 中的元素 arr2[i] 各不相同 
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中 
// 
// Related Topics 排序 数组
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	array := relativeSortArray(arr1, arr2)
	fmt.Println(array)
}

//leetcode submit region begin(Prohibit modification and deletion)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	// 桶：个人写法
	n := len(arr2)
	bucket, diffs := make([]int, n), make([]int, 0)
	memo := make(map[int]int)
	for i, v := range arr2 { // O(n)
		memo[v] = i
	}
	for _, v := range arr1 { // O(m)
		if i, ok := memo[v]; ok {
			bucket[i]++
		} else {
			diffs = append(diffs, v)
		}
	}
	ans, i := make([]int, len(arr1)), 0
	for j, count := range bucket { // O(m-n)
		for k := 0; k < count; k++ {
			ans[i], i = arr2[j], i+1
		}
	}
	sort.Ints(diffs)     // O((m-n)long (m-n))
	copy(ans[i:], diffs) // O(m-n)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
