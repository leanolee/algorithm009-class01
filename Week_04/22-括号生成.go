//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 
//
// 示例： 
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
// 
// Related Topics 字符串 回溯算法
package main

import "fmt"

func main() {
	n := 3
	parenthesis := generateParenthesis(n)
	fmt.Println(parenthesis)
}

type Parenthesis struct {
	str    []uint8
	nLeft  int
	nRight int
}

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	// BFS

	// DFS
	parenthesis := make([]string, 0)
	parenthesisRecursion(n, &parenthesis, nil, 0, 0)
	return parenthesis
}

func parenthesisRecursion(n int, res *[]string, path []uint8, left int, right int) {
	if left == n && right == n {
		*res = append(*res, string(path))
		return
	}
	if left < n {
		parenthesisRecursion(n, res, append(append([]uint8{}, path...), '('), left+1, right)
	}
	if right < left {
		parenthesisRecursion(n, res, append(append([]uint8{}, path...), ')'), left, right+1)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
