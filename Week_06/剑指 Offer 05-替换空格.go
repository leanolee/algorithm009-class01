//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。 
//
// 
//
// 示例 1： 
//
// 输入：s = "We are happy."
//输出："We%20are%20happy." 
//
// 
//
// 限制： 
//
// 0 <= s 的长度 <= 10000 
//
package main

import "fmt"

func main() {
	s := "We are happy."
	space := replaceSpace(s)
	fmt.Println(space)
}

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	sByte := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			sByte = append(sByte, "%20"...)
		} else {
			sByte = append(sByte, s[i])
		}
	}
	return string(sByte)
}

//leetcode submit region end(Prohibit modification and deletion)
