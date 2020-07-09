//你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方
//法，生成跳水板所有可能的长度。 
// 返回的长度需要从小到大排列。 
// 示例： 
// 输入：
//shorter = 1
//longer = 2
//k = 3
//输出： {3,4,5,6}
// 
// 提示： 
// 
// 0 < shorter <= longer 
// 0 <= k <= 100000 
// 
// Related Topics 递归 记忆化
package main

import "fmt"

func main() {
	shorter := 1
	longer := 2
	k := 3
	board := divingBoard(shorter, longer, k)
	fmt.Println(board)
}

//leetcode submit region begin(Prohibit modification and deletion)
func divingBoard(shorter int, longer int, k int) []int {
	// dp
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	ans := make([]int, k+1)
	off, pre := longer-shorter, shorter*k
	for i := 0; i <= k; i++ {
		ans[i], pre = pre, pre+off
	}
	return ans

	// dfs：超时
	//if k <= 0 {
	//	return nil
	//}
	//board := make([]int, 0)
	//dfs(shorter, longer, k, &board, 0, false)
	//return board
}

func dfs(shorter int, longer int, k int, ans *[]int, sum int, short bool) {
	if k == 0 {
		*ans = append(*ans, sum)
		return
	}
	if !short {
		dfs(shorter, longer, k-1, ans, sum+shorter, false)
	}
	dfs(shorter, longer, k-1, ans, sum+longer, true)
}

//leetcode submit region end(Prohibit modification and deletion)
