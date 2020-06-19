//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换
//需遵循如下规则： 
//
// 
// 每次转换只能改变一个字母。 
// 转换后得到的单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回一个空列表。 
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
//输出:
//[
//  ["hit","hot","dot","dog","cog"],
//  ["hit","hot","lot","log","cog"]
//]
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。 
// Related Topics 广度优先搜索 数组 字符串 回溯算法
package main

import (
	"fmt"
	"math"
)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ladders := findLadders(beginWord, endWord, wordList)
	fmt.Println(ladders)

	//test := math.MaxInt64
	//fmt.Println(test)
	//fmt.Println(test + 1)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// BFS
	n := len(wordList)
	memo := make(map[string]int)
	for i, word := range wordList {
		memo[word] = i
	}
	if _, ok := memo[beginWord]; !ok { // 将beginWord
		wordList = append(wordList, beginWord)
		memo[beginWord] = n
		n++
	}
	if _, ok := memo[endWord]; n == 0 || !ok {
		return nil
	}
	findMap := make([][]int, n) // 存储和当前位置的 word 能产生变换的其他 word 的索引
	for i, word := range wordList {
		for j := i + 1; j < n; j++ {
			if canLadders(word, wordList[j]) {
				findMap[i] = append(findMap[i], j)
				findMap[j] = append(findMap[j], i)
			}
		}
	}
	//fmt.Println(findMap)
	queue := [][]int{{memo[beginWord]}} // 记录所有变换的路径
	findLadders := make([][]string, 0)
	cost := make([]int, n) // 记录 worldList 同索引的字符串，变为 endWord 需要的最小次数
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[memo[beginWord]] = 0
	for i := 0; i < len(queue); i++ {
		curr := queue[i]
		lastIdx := curr[len(curr)-1] // 最后一个word的索引
		if lastIdx == memo[endWord] {
			temp := make([]string, len(curr))
			for j, idx := range curr {
				temp[j] = wordList[idx]
			}
			findLadders = append(findLadders, temp)
		} else {
			for _, nextIdx := range findMap[lastIdx] {
				if cost[lastIdx]+1 <= cost[nextIdx] {
					cost[nextIdx] = cost[lastIdx] + 1
					plus := make([]int, len(curr))
					copy(plus, curr)
					queue = append(queue, append(plus, nextIdx))
				}
			}
		}
	}
	return findLadders
}
func canLadders(begin, end string) bool {
	for i := 0; i < len(begin); i++ {
		if begin[i] != end[i] {
			return begin[i+1:] == end[i+1:]
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
