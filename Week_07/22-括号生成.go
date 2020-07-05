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

type P struct {
	l int
	r int
	s []byte
}

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	// BFS
	ps := []P{{0, 0, nil}}
	ans := make([]string, 0)
	for i := n << 1; i > 0; i-- {
		m := len(ps)
		for j := 0; j < m; j++ {
			l, r, s := ps[j].l, ps[j].r, ps[j].s
			if l == n {
				ps[j].s, ps[j].r = append(s, ')'), r+1
				continue
			}
			if l == r {
				ps[j].s, ps[j].l = append(s, '('), l+1
				continue
			}
			ps[j].s, ps[j].l = append(s, '('), l+1
			ps = append(ps, P{l, r + 1, append(append([]byte{}, s...), ')')})
		}
	}
	for _, v := range ps {
		ans = append(ans, string(v.s))
	}
	return ans

	// dfs
	//ans := make([]string, 0)
	//generateParenthesisRecursion(n, &ans, nil, 0, 0)
	//return ans
}

func generateParenthesisRecursion(n int, ans *[]string, path []byte, l int, r int) {
	if len(path) == n<<1 {
		*ans = append(*ans, string(path))
		return
	}
	next := make([]byte, len(path))
	copy(next, path)
	if l < n {
		next = append(next, '(')
		generateParenthesisRecursion(n, ans, next, l+1, r)
		next = next[:len(next)-1]
	}
	if r < l { // 剪枝
		generateParenthesisRecursion(n, ans, append(next, ')'), l, r+1)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
