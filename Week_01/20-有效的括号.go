//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串
package main

import "fmt"

func main() {
	//s := "[(){}[]({})]"
	s := "()[]{}"
	valid := isValid(s)
	fmt.Println(valid)
}

/*
替换思路：每次存入栈中，存括号对应的右半部分，这样遇到右括号时，可以直接比较

暴力法：
	不断循环遍历 replace 匹配的括号为 ""，直到一次遍历循环没有发现相邻的两个括号匹配
	如果最后剩下 ""，则返回true
开心消消乐：O(n^2),O(n)
	个人方法
	1.将目标字符串放入 stack 字符数组，将用一个指针 index 记录下一个要消灭的 左括号 的 右边 那个索引
	2.一次循环：遍历原字符串
		当是左括号时，index++
		当是右括号时，如果 stack[index-1] 与 当前字符s[i] 互为括号，index--；否则失败
	3.按 index 标记截取 stack[:index-1] + stack[index+1:]
开心消消乐：优化，O(n),O(n)
	1.创建用于存放左括号的数组 left[]，用一个变量记录存入的数量
	2.一次循环：遍历原字符串
		当是左括号，则存入数组中，且存入的是 左括号对应的右括号，count++
		当是右括号，如果 left[count] == s[i]，则count--；否则返回失败
			并且注意，count此时不能为 负值
	3.最后返回count == 0
Stack：O(n),O(n)
	1.创建 stack
	2.一次循环：遍历原字符串
		如果是左括号，压入 stack 中
		如果是右括号，若与栈顶括号匹配，弹出栈顶元素；否则返回失败
	3.返回 stack.len == 0
*/
//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	// 个人方法
	//n := len(s)
	//lefts := make([]byte, n)
	//j := 0
	//for i := 0; i < n; i++ {
	//	switch s[i] {
	//	case '(':
	//		lefts[j] = ')'
	//		j++
	//	case '[':
	//		lefts[j] = ']'
	//		j++
	//	case '{':
	//		lefts[j] = '}'
	//		j++
	//	default:
	//		if j == 0 || lefts[j-1] != s[i] {
	//			return false
	//		}
	//		j--
	//	}
	//}
	//return j == 0

	// 栈
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) > 0 && (s[i] == ')' && stack[len(stack)-1] == '(' || s[i] == ']' && stack[len(stack)-1] == '[' || s[i] == '}' && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0

	// 开心消消乐：双百
	//stack := []byte(s)
	//for i, index := 0, 0; i < len(s); i++ {
	//	if s[i] == '(' || s[i] == '[' || s[i] == '{' {
	//		index++
	//		continue
	//	}
	//	if index > 0 && (s[i] == ')' && stack[index-1] == '(' || s[i] == ']' && stack[index-1] == '[' || s[i] == '}' && stack[index-1] == '{') {
	//		stack = append(stack[:index-1], stack[index+1:]...)
	//		index--
	//	} else {
	//		return false
	//	}
	//}
	//return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
