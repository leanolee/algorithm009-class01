//给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。 
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。 
//
// 示例1: 
//
// 输入: pattern = "abba", str = "dog cat cat dog"
//输出: true 
//
// 示例 2: 
//
// 输入:pattern = "abba", str = "dog cat cat fish"
//输出: false 
//
// 示例 3: 
//
// 输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false 
//
// 示例 4: 
//
// 输入: pattern = "abba", str = "dog dog dog dog"
//输出: false 
//
// 说明: 
//你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。 
// Related Topics 哈希表
package main

import (
	"fmt"
	"strings"
)

func main() {
	//pattern := "abba"
	//str := "dog dog dog dog"
	//pattern := "abba"
	//str := "dog cat cat dog"
	pattern := "abba"
	str := "fish whoops helloworld fish"
	b := wordPattern(pattern, str)
	fmt.Println(b)
}

/*
暴力法：O(n^2),O(1)
	1.切割字符串为数组，长度不同，返回false
	2.一层循环：从头至尾，分别锁定 pattern 和 数组 的一个元素
	3.二层循环：（pattern、数组 分别从头至尾与锁定元素比较，一个同，另外一个也要相同；一个不同，另外一个也要不同
双hash映射：O(n),O(n)
	1.切割，长度不同，返回
	2.一次循环：两个对象分别放入一个映射中 c->s、s->c，互相映射
	3.二次循环：取对应的映射 c != c || c != c；返回false
双hash映射优化：O(n),O(n)
	1.切割，长度不同，返回
	2.双映射初始化，循环遍历 pattern(c)、字符串数组(s)：
		c 存在于映射中，取出 s，与 s 数组对比，不同，返回false
		c 不存在于映射中，看 s 映射中是否有c，有，返回false
			将 c、s分别存入映射
*/
//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, str string) bool {
	// 双map
	//split := strings.Split(str, " ")
	//if len(split) != len(pattern) {
	//	return false
	//}
	//chars := make(map[int32]string)
	//strs := make(map[string]int32)
	//for k, v := range pattern {
	//	if c, ok := chars[v]; ok {
	//		if c != split[k] {
	//			return false
	//		}
	//	} else {
	//		if _, ok := strs[split[k]]; ok {
	//			return false
	//		}
	//		chars[v] = split[k]
	//		strs[split[k]] = v
	//	}
	//}
	//return true

	// 个人方法
	split := strings.Split(str, " ")
	if len(pattern) != len(split) {
		return false
	}
	memoChar := make(map[byte]string)
	memoStr := make(map[string]byte)
	for i := 0; i < len(split); i++ {
		memoChar[pattern[i]] = split[i]
		memoStr[split[i]] = pattern[i]
	}
	for i := 0; i < len(split); i++ {
		if memoChar[pattern[i]] != split[i] || memoStr[split[i]] != pattern[i] {
			return false
		}
	}
	return true

	// 暴力法
	//split := strings.Split(str, " ")
	//if len(pattern) != len(split) {
	//	return false
	//}
	//for i := 0; i < len(pattern); i++ {
	//	for j := i + 1; j < len(pattern); j++ {
	//		if pattern[i] == pattern[j] && split[i] != split[j] || pattern[i] != pattern[j] && split[i] == split[j] {
	//			return false
	//		}
	//	}
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
