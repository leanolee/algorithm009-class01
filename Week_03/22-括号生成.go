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
	str    string
	nLeft  int
	nRight int
}

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	// dp
	/*
		参考：https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/
		n：dp[i] = "(" + dp[j] + ")" + dp[i- j - 1] , j = 0, 1, ..., i - 1
			n-1：所包含的情况，外面增加 (...)
			p+q=n-1：(p)q，
			上面两者相加
	*/
	// 迭代
	parenthesis := []Parenthesis{{"", 0, 0}}
	for i := 0; i < n<<1; i++ {
		length := len(parenthesis)
		for j := 0; j < length; j++ {
			p := parenthesis[j]
			if p.nLeft == p.nRight { // 左右括号数相等
				parenthesis[j].str += "("
				parenthesis[j].nLeft++
				continue
			}
			if p.nLeft == n { // 左括号已满
				parenthesis[j].str += ")"
				parenthesis[j].nRight++
				continue
			}
			parenthesis[j].str += "(" // n > left > right
			parenthesis[j].nLeft++
			parenthesis = append(parenthesis, Parenthesis{p.str + ")", p.nLeft, p.nRight + 1})
		}
	}
	res := make([]string, len(parenthesis))
	for i, p := range parenthesis {
		res[i] = p.str
	}
	return res

	// 回溯
	//parenthesis := make([]string, 0)
	//generateParenthesisRecursion(0, 0, "", &parenthesis, n)
	//return parenthesis

	// 递归
	//parenthesis := make([]string, 0)
	//generateParenthesisRecursion(0, 0, "", &parenthesis, n)
	//return parenthesis
}

func generateParenthesisRecursion(left int, right int, s string, parenthesis *[]string, n int) {
	if left == n && right == n {
		*parenthesis = append(*parenthesis, s)
		return
	}
	if left < n { // 剪枝
		s += "("
		generateParenthesisRecursion(left+1, right, s, parenthesis, n)
		s = s[:len(s)-1] // 回溯
	}
	if left > right {
		s += ")"
		generateParenthesisRecursion(left, right+1, s, parenthesis, n)
		s = s[:len(s)-1]
	}

	//if left == n && right == n {
	//	*parenthesis = append(*parenthesis, s)
	//	return
	//}
	//if left < n {
	//	generateParenthesisRecursion(left+1, right, s+"(", parenthesis, n)
	//}
	//if left > right {
	//	generateParenthesisRecursion(left, right+1, s+")", parenthesis, n)
	//}	// 这里不需要回溯，因为 s 的值是在参数中修改，并没有修改本函数中 s 的值
}

//leetcode submit region end(Prohibit modification and deletion)
