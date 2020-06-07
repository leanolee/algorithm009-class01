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

import (
	"fmt"
)

func main() {
	s := "We are happy."
	space := replaceSpace(s)
	fmt.Println(space)
}

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	// rune
	str := make([]rune, 0)
	for _, c := range s {
		if c == 32 {
			str = append(str, 37, 50, 48)
		} else {
			str = append(str, c)
		}
	}
	return string(str)

	// Builder
	//var str strings.Builder
	//for _, c := range s {
	//	if c == ' ' {
	//		str.WriteString("%20")
	//	} else {
	//		str.WriteRune(c)
	//	}
	//}
	//return str.String()

	// 个人写法
	//space := make([]byte, 0)
	//for _, c := range s {
	//	if c == ' ' {
	//		space = append(space, "%20"...)
	//	} else {
	//		space = append(space, byte(c))
	//	}
	//}
	//return string(space)
}

//leetcode submit region end(Prohibit modification and deletion)
