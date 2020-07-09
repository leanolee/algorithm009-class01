//给出一个区间的集合，请合并所有重叠的区间。 
//
// 示例 1: 
//
// 输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 
//
// 示例 2: 
//
// 输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。 
// Related Topics 排序 数组
package main

import "fmt"

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	i := merge(intervals)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	// 归并：个人写法
	if len(intervals) == 0 {
		return nil
	}
	return mergeSort(intervals, 0, len(intervals)-1)
}

func mergeSort(arr [][]int, l int, r int) [][]int {
	if l >= r {
		return [][]int{arr[r]}
	}
	mid := (l + r) >> 1
	return mergeArr(mergeSort(arr, l, mid), mergeSort(arr, mid+1, r))
}

func mergeArr(left [][]int, right [][]int) [][]int {
	/*
		写法一：先排序，再合并区间
		写法二：边排序边合并区间，代码多点，但是更效率
	*/
	// 写法二
	m, n := len(left), len(right)
	i, j, ans := 0, 0, make([][]int, 0)
	var temp []int
	for i < m && j < n {
		if left[i][0] < right[j][0] {
			temp, i = left[i], i+1
		} else {
			temp, j = right[j], j+1
		}
		if len(ans) == 0 {
			ans = append(ans, temp)
			continue
		}
		mergeIntervals(&ans, len(ans)-1, temp)
	}
	for ; i < m; i++ {
		mergeIntervals(&ans, len(ans)-1, left[i])
	}
	for ; j < n; j++ {
		mergeIntervals(&ans, len(ans)-1, right[j])
	}
	return ans

	// 写法一
	//m, n := len(left), len(right)
	//i, j, k, temp := 0, 0, 0, make([][]int, m+n)
	//for ; i < m && j < n; k++ {
	//	if left[i][0] < right[j][0] {
	//		temp[k], i = left[i], i+1
	//	} else {
	//		temp[k], j = right[j], j+1
	//	}
	//
	//}
	//copy(temp[k:], left[i:])
	//copy(temp[k:], right[j:])
	//ans := [][]int{temp[0]}
	////ans := make([][]int, 0)
	//for idx, curr := 1, 0; idx < len(temp); idx++ {
	//	if ans[curr][1] < temp[idx][0] {
	//		ans, curr = append(ans, temp[idx]), curr+1
	//	} else if ans[curr][1] < temp[idx][1] {
	//		ans[curr][1] = temp[idx][1]
	//	}
	//}
	//return ans
}

func mergeIntervals(ans *[][]int, last int, temp []int) {
	end := (*ans)[last][1]
	if end < temp[0] {
		*ans = append(*ans, temp)
	} else if end < temp[1] {
		(*ans)[last][1] = temp[1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
