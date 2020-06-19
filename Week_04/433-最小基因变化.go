//一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。 
//
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。 
//
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。 
//
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。 
//
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变
//化次数。如果无法实现目标变化，请返回 -1。 
//
// 注意: 
//
// 
// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。 
// 所有的目标基因序列必须是合法的。 
// 假定起始基因序列与目标基因序列是不一样的。 
// 
//
// 示例 1: 
//
// 
//start: "AACCGGTT"
//end:   "AACCGGTA"
//bank: ["AACCGGTA"]
//
//返回值: 1
// 
//
// 示例 2: 
//
// 
//start: "AACCGGTT"
//end:   "AAACGGTA"
//bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//
//返回值: 2
// 
//
// 示例 3: 
//
// 
//start: "AAAAACCC"
//end:   "AACCCCCC"
//bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//
//返回值: 3
// 
//
package main

import (
	"fmt"
)

func main() {
	//start := "AACCGGTT"
	//end := "AAACGGTA"
	//bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	//start := "AACCGGTT"
	//end := "AACCGGTA"
	//bank := []string{}
	start := "AAAAAAAA"
	end := "CCCCCCCC"
	bank := []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"}
	mutation := minMutation(start, end, bank)
	fmt.Println(mutation)
}

type Mutation struct {
	gene      []uint8
	childrens []*Mutation
}

/*
参考：leetcode-17、BFS、DFS也可以

很麻烦的方法：
	如果这个变化序列需要长期使用，可以建立对应层序树状结构
	但是需要是可以往回指的：.parent = *(父节点的指针)
		或者子节点数组中第一个元素默认为 父节点
*/
//leetcode submit region begin(Prohibit modification and deletion)
func minMutation(start string, end string, bank []string) int {
	// 双向BFS：双百
	idx := -1
	for i, b := range bank {
		if end == b {
			idx = i
			break
		}
	}
	if len(bank) == 0 || idx == -1 {
		return -1
	}
	canMutation := func(curr, to string) bool {
		count := 0
		for i := 0; i < len(curr); i++ {
			if curr[i] != to[i] {
				if count == 1 {
					return false
				}
				count++
			}
		}
		return true
	}
	count := 0
	mutation := func(queue *[]string, currV, otherV []bool) int {
		n := len(*queue)
		for i := 0; i < n; i++ {
			curr := (*queue)[i]
			for i, b := range bank {
				if !currV[i] && canMutation(curr, b) {
					if otherV[i] {
						return 0
					}
					currV[i] = true
					*queue = append(*queue, b)
				}
			}
		}
		*queue = (*queue)[n:]
		return -1
	}
	qStart, sVisited, qEnd, eVisited := []string{start}, make([]bool, len(bank)), []string{end}, make([]bool, len(bank))
	eVisited[idx] = true
	for len(qStart) > 0 && len(qEnd) > 0 {
		if can := mutation(&qStart, sVisited, eVisited); can == 0 {
			return (count << 1) + 1
		}
		if can := mutation(&qEnd, eVisited, sVisited); can == 0 {
			return (count << 1) + 2
		}
		count++
	}
	return -1

	// BFS：双百
	//canMutation := func(curr, to string) bool {
	//	count := 0
	//	for i := 0; i < len(curr); i++ {
	//		if curr[i] != to[i] {
	//			if count == 1 {
	//				return false
	//			}
	//			count++
	//		}
	//	}
	//	return true
	//}
	//used := make([]bool, len(bank))
	//queue := []string{start}
	//for i, min := 0, 1; len(queue) > i; min++ { // queue的新写法
	//	n := len(queue)
	//	for ; i < n; i++ {
	//		curr := queue[i]
	//		for i, mu := range bank {
	//			if canMutation(curr, mu) && !used[i] {
	//				if mu == end {
	//					return min
	//				}
	//				used[i] = true
	//				queue = append(queue, mu)
	//			}
	//		}
	//	}
	//}
	//return -1

	// DFS：双百
	//minMutation := len(bank) + 1
	//isUsed := make([]bool, len(bank))
	//canMutation := func(curr, mu string) bool {
	//	count := 0
	//	for i, c := range curr {
	//		if uint8(c) != mu[i] {
	//			if count == 1 {
	//				return false
	//			}
	//			count++
	//		}
	//	}
	//	return true
	//}
	//var dfs func(curr string, level int, used []bool)
	//dfs = func(curr string, level int, used []bool) {
	//	for i, muTo := range bank {
	//		if used[i] || !canMutation(curr, muTo) {
	//			continue
	//		}
	//		if muTo == end {
	//			if level+1 < minMutation {
	//				minMutation = level + 1
	//			}
	//			return
	//		}
	//		used[i] = true
	//		dfs(muTo, level+1, used)
	//		used[i] = false
	//	}
	//}
	//dfs(start, 0, isUsed)
	//if minMutation <= len(bank) {
	//	return minMutation
	//}
	//return -1

	// 个人写法：BFS+hash
	//memo := make(map[string]bool)
	//for _, s := range bank {
	//	memo[s] = true
	//}
	//if _, ok := memo[end]; !ok { // 先检查 bank 中是否存在end
	//	return -1
	//}
	//delete(memo, end)
	//diff := 0
	//targets := map[string]int{end: diff}
	//for len(targets) > 0 { // 如果 ==0，说明相差 >1
	//	diff++
	//	for tar, _ := range targets { // 比较和 start 是否相差 1
	//		if getDiff(tar, start) == 1 {
	//			return diff
	//		}
	//	}
	//	for tar, d := range targets { // 开始下一轮的准备
	//		if d > diff { // 重点1：用diff标记是哪一轮
	//			continue
	//		}
	//		for k, _ := range memo {
	//			if getDiff(tar, k) == 1 { // 取出相差1个的
	//				targets[k] = diff + 1 // 下一轮的数据，同时用 diff+1 标记
	//				delete(memo, k)       // 已添加到下一轮的数据，从 bank 中移除
	//			}
	//		}
	//		delete(targets, tar) // 重点2：targets 中数据有两个用途：1.比较和 start 是否差 1；2.从 bank 中找出与 targets 相差 1 的元素，备用下一轮比较
	//	}
	//}
	//return -1
}

func getDiff(tar string, another string) int {
	count := 0
	for i := 0; i < len(tar); i++ {
		if tar[i] != another[i] {
			count++
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
