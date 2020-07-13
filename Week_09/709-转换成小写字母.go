//实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。 
//
// 
//
// 示例 1： 
//
// 
//输入: "Hello"
//输出: "hello" 
//
// 示例 2： 
//
// 
//输入: "here"
//输出: "here" 
//
// 示例 3： 
//
// 
//输入: "LOVELY"
//输出: "lovely"
// 
// Related Topics 字符串
package main

import "fmt"

func main() {
	str := "LOVELY"
	lowerCase := toLowerCase(str)
	fmt.Println(lowerCase)
}

//leetcode submit region begin(Prohibit modification and deletion)
func toLowerCase(str string) string {
	chars := []byte(str)
	for i, c := range chars {
		if c >= 65 && c <= 90 {
			chars[i] = c + 32
		}
	}
	return string(chars)
}

//leetcode submit region end(Prohibit modification and deletion)
