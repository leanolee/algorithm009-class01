//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
// 
//
// 
// 每次转换只能改变一个字母。 
// 转换过程中的中间单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回 0。 
// 所有单词具有相同的长度。 
// 所有单词只由小写字母组成。 
// 字典中不存在重复的单词。 
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。 
// 
//
// 示例 1: 
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。 
// Related Topics 广度优先搜索
package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	length := ladderLength(beginWord, endWord, wordList)
	fmt.Println(length)
}

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// BFS
	n := len(beginWord)
	wordTree := make(map[string][]string) // 存储映射关系
	for _, str := range wordList {
		for i := 0; i < n; i++ {
			key := str[:i] + "_" + str[i+1:]
			wordTree[key] = append(wordTree[key], str)
		}
	}
	visited := make(map[string]bool) // 已经访问过的 wordList 中的字符串
	targets := make(map[string]int)  // 存储目标字符串
	level := 1
	targets[beginWord] = level
	temp := make(map[string]int) // 由于map遍历的问题，需要个中转的map。将下一次的字符串装入temp，先清空 targets，再将temp倒入targets
	for len(targets) > 0 {
		for kStr, vLevel := range targets {
			for i := 0; i < n; i++ {
				for _, s := range wordTree[kStr[:i]+"_"+kStr[i+1:]] {
					if s == endWord {
						return vLevel + 1
					} else if !visited[s] {
						visited[s] = true
						temp[s] = level + 1
					}
				}
			}
		}
		level++
		targets, temp = temp, make(map[string]int)
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
