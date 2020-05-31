//有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"
//，"(())()" 和 "(()(()))" 都是有效的括号字符串。 
//
// 如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。 
//
// 给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。 
//
// 对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。 
//
// 
//
// 示例 1： 
//
// 输入："(()())(())"
//输出："()()()"
//解释：
//输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
//删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。 
//
// 示例 2： 
//
// 输入："(()())(())(()(()))"
//输出："()()()()(())"
//解释：
//输入字符串为 "(()())(())(()(()))"，原语化分解得到 "(()())" + "(())" + "(()(()))"，
//删除每个部分中的最外层括号后得到 "()()" + "()" + "()(())" = "()()()()(())"。
// 
//
// 示例 3： 
//
// 输入："()()"
//输出：""
//解释：
//输入字符串为 "()()"，原语化分解得到 "()" + "()"，
//删除每个部分中的最外层括号后得到 "" + "" = ""。
// 
//
// 
//
// 提示： 
//
// 
// S.length <= 10000 
// S[i] 为 "(" 或 ")" 
// S 是一个有效括号字符串 
// 
// Related Topics 栈
package main

import (
	"fmt"
)

func main() {
	S := "(()())(())"
	parentheses := removeOuterParentheses(S)
	fmt.Println(parentheses)
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeOuterParentheses(S string) string {
	// 数组
	//stack := make([]uint8, 0)
	//leftCount := 1
	//for i := 1; i < len(S)-1; i++ {
	//	if S[i] == '(' {
	//		leftCount++
	//	} else {
	//		leftCount--
	//		if leftCount == 0 {
	//			leftCount++
	//			i++
	//			continue
	//		}
	//	}
	//	stack = append(stack, S[i])
	//}
	//return string(stack)

	// 个人方法：字符串拼接
	j := -1
	leftCount := 0
	var parentheses string
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			leftCount++
		} else {
			leftCount--
		}
		if leftCount == 0 {
			parentheses += S[j+2 : i]
			j = i
		}
	}
	return parentheses
}

//leetcode submit region end(Prohibit modification and deletion)
