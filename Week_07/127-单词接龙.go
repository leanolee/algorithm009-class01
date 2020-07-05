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

import (
	"container/list"
	"fmt"
)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	length := ladderLength(beginWord, endWord, wordList)
	fmt.Println(length)
}

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 双向BFS

	// BFS
	count, n := 1, len(wordList)
	queue, visited := list.New(), make([]bool, n)
	queue.PushBack(beginWord)
	getLadder := func(begin, tar string) bool {
		for i := 0; i < len(begin); i++ {
			if begin[i] != tar[i] {
				return begin[i+1:] == tar[i+1:]
			}
		}
		return false
	}
	for queue.Len() > 0 {
		count++
		n := queue.Len()
		for i := 0; i < n; i++ {
			begin := queue.Remove(queue.Front()).(string)
			for i, word := range wordList {
				if !visited[i] && getLadder(begin, word) {
					if word == endWord {
						return count
					}
					visited[i] = true
					queue.PushBack(word)
				}
			}
		}
	}
	return 0

	// DFS
}

//leetcode submit region end(Prohibit modification and deletion)
