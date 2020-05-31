//统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。 
//
// 请注意，你可以假定字符串里不包括任何不可打印的字符。 
//
// 示例: 
//
// 输入: "Hello, my name is John"
//输出: 5
//解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
// 
// Related Topics 字符串
package main

import (
	"fmt"
)

func main() {
	s := "Hello, my name is John"
	segments := countSegments(s)
	fmt.Println(segments)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countSegments(s string) int {
	count := 0
	segment := false
	for _, c := range s {
		if c != ' ' {
			if !segment {
				count++
				segment = true
			}
		} else {
			segment = false
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
