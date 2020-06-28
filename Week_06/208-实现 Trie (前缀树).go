//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。 
//
// 示例: 
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");   
//trie.search("app");     // 返回 true 
//
// 说明: 
//
// 
// 你可以假设所有的输入都是由小写字母 a-z 构成的。 
// 保证所有输入均为非空字符串。 
// 
// Related Topics 设计 字典树
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	next [26]*Trie
	word string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		idx := c - 'a'
		if curr.next[idx] == nil {
			trie := &Trie{}
			curr.next[idx], curr = trie, trie
		} else {
			curr = curr.next[idx]
		}
	}
	curr.word = word
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		idx := c - 'a'
		if curr.next[idx] == nil {
			return false
		}
		curr = curr.next[idx]
	}
	return curr.word != "" // 结构体写为：word string，就要这样判断
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		idx := c - 'a'
		if curr.next[idx] == nil {
			return false
		}
		curr = curr.next[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
