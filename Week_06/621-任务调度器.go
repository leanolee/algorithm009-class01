//给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务
//都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。 
//
// 然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。 
//
// 你需要计算完成所有任务所需要的最短时间。 
//
// 
//
// 示例 ： 
//
// 输入：tasks = ["A","A","A","B","B","B"], n = 2
//输出：8
//解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
//     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。 
//
// 
//
// 提示： 
//
// 
// 任务的总个数为 [1, 10000]。 
// n 的取值范围为 [0, 100]。 
// 
// Related Topics 贪心算法 队列 数组
package main

import (
	"fmt"
	"sort"
)

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2
	interval := leastInterval(tasks, n)
	fmt.Println(interval)
}

//leetcode submit region begin(Prohibit modification and deletion)
func leastInterval(tasks []byte, n int) int {
	// 设计：为什么这么牛逼？
	memo := make([]int, 26)
	for _, t := range tasks {
		memo[t-'A']++
	}
	sort.Ints(memo)
	empty := (memo[25] - 1) * n
	for i := 24; i >= 0 && memo[i] != 0; i-- {
		if memo[i] == memo[25] {
			empty++
		}
		empty -= memo[i]
	}
	if empty > 0 {
		return empty + len(tasks)
	}
	return len(tasks)

	// PriorityQueue：

	// 排序
	//memo := make([]int, 26)
	//for _, t := range tasks {
	//	memo[t-'A']++
	//}
	//sort.Ints(memo)
	//plus := 0
	//for {
	//	count := 0
	//	for i := 25; i >= 0 && count <= n; i-- {
	//		if memo[i] == 0 {
	//			break
	//		}
	//		memo[i]--
	//		count++
	//	}
	//	sort.Ints(memo)
	//	if memo[25] == 0 {
	//		break
	//	}
	//	plus += n + 1 - count
	//}
	//return len(tasks) + plus
}

//leetcode submit region end(Prohibit modification and deletion)
