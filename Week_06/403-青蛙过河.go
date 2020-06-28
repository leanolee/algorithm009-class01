//一只青蛙想要过河。 假定河流被等分为 x 个单元格，并且在每一个单元格内都有可能放有一石子（也有可能没有）。 青蛙可以跳上石头，但是不可以跳入水中。 
//
// 给定石子的位置列表（用单元格序号升序表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一个石子上）。 开始时， 青蛙默认已站在第一个石子上，并可以
//假定它第一步只能跳跃一个单位（即只能从单元格1跳至单元格2）。 
//
// 如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。 
//
//
// 请注意： 
//
// 
// 石子的数量 ≥ 2 且 < 1100； 
// 每一个石子的位置序号都是一个非负整数，且其 < 231； 
// 第一个石子的位置永远是0。 
// 
//
// 示例 1: 
//
// 
//[0,1,3,5,6,8,12,17]
//
//总共有8个石子。
//第一个石子处于序号为0的单元格的位置, 第二个石子处于序号为1的单元格的位置,
//第三个石子在序号为3的单元格的位置， 以此定义整个数组...
//最后一个石子处于序号为17的单元格的位置。
//
//返回 true。即青蛙可以成功过河，按照如下方案跳跃： 
//跳1个单位到第2块石子, 然后跳2个单位到第3块石子, 接着 
//跳2个单位到第4块石子, 然后跳3个单位到第6块石子, 
//跳4个单位到第7块石子, 最后，跳5个单位到第8个石子（即最后一块石子）。
// 
//
// 示例 2: 
//
// 
//[0,1,2,3,4,8,9,11]
//
//返回 false。青蛙没有办法过河。 
//这是因为第5和第6个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。
// 
// Related Topics 动态规划
package main

import "fmt"

func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	//stones := []int{0, 1, 2, 3, 4, 8, 9, 11} // false
	//stones := []int{0, 1, 3, 6, 10, 15, 16, 21}
	cross := canCross(stones)
	fmt.Println(cross)
}

//leetcode submit region begin(Prohibit modification and deletion)
func canCross(stones []int) bool {
	// dp：
	n := len(stones)
	memo := make(map[int]map[int]bool)
	for _, s := range stones {
		memo[s] = make(map[int]bool)
	}
	memo[0] = map[int]bool{0: true}
	for i := 0; i < n; i++ {
		for k, _ := range memo[stones[i]] {
			for step := k - 1; step <= k+1; step++ {
				key := stones[i] + step
				if step == 0 || memo[key] == nil {
					continue
				}
				memo[key][step] = true
			}
			//fmt.Println(i, "->", memo)
		}
	}
	return len(memo[stones[n-1]]) > 0

	// 记忆化 + 二分：查找下一步是否有石头，不过二分太偏前，有必要？

	// 递归：记忆化
	//memo := make(map[int]bool) // step*n+idx
	//if stones[1] != 1 {
	//	return false
	//}
	//n := len(stones)
	//return crossRecursion(stones, memo, n, 1, 1)

	// 递归：超时
	//if stones[1] != 1 {
	//	return false
	//}
	//return canCrossRecursion(stones, 1, 1)
}

func crossRecursion(stones []int, memo map[int]bool, n int, idx int, step int) bool {
	if idx == len(stones)-1 {
		return true
	}
	key := step*n + idx
	if b, ok := memo[key]; ok {
		return b
	}
	can := false
	for i := idx + 1; i < len(stones); i++ {
		nextStep := stones[i] - stones[idx]
		if step-1 > nextStep {
			continue
		} else if nextStep > step+1 {
			break
		}
		next := crossRecursion(stones, memo, n, i, nextStep)
		memo[nextStep*n+i] = next
		can = can || next
	}
	memo[key] = can
	return can
}

func canCrossRecursion(stones []int, idx int, step int) bool { // idx：目标，step：当前可跳步数
	if idx == len(stones)-1 {
		return true
	}
	can := false
	for i := idx + 1; i < len(stones); i++ {
		nextStep := stones[i] - stones[idx]
		if step-1 > nextStep {
			continue
		} else if nextStep > step+1 {
			break
		}
		can = can || canCrossRecursion(stones, i, nextStep)
	}
	return can
}

//leetcode submit region end(Prohibit modification and deletion)
