//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？ 
//
// 示例: 
//
// 输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3 
// Related Topics 树 动态规划
package main

import "fmt"

func main() {
	n := 5 // [1 1 2 5 14 42 132 429 1430 4862 16796 58786 208012 742900 2674440 9694845 35357670 129644790 477638700 1767263190 6564120420 24466267020 91482563640 343059613650 1289904147324 4861946401452 18367353072152 69533550916004 263747951750360 1002242216651368 3814986502092304]
	trees := numTrees(n)
	fmt.Println(trees)
}

/*
记忆化递归：
思路：观察示例 n = 3，答案是 [1,2,3] 的全排列的子集
	1.所以是否可以通过dp来求解呢？
		猜想是可以的，但是一时没有想出来
	2.那么通过递归来求解应该是可以的吧？
		如 n = 3，BST中有且仅有 1,2,3 这3个节点
			1.根节点选为 1：那么 2、3 节点必然为右节点（因为 2、3 都大于 1）
			2.可以递推，当根节点为 1 时，左子树节点数为 0 个，右子树节点数为 2 个
				此时BST的情况可以有两种
			3.显然，可以将上述递推推广为根节点为 2、3，或任意节点
				BST的总数 = 左子树的情况数 * 右子树的情况数
			4.重要：由于 [1,2,3] 数组是有序的，所以不管根节点选择是哪个，将原数组（[1,2,3]）分为两部分后 左数组 和 右数组 都是分别有序的（且不重复）
				加上BST的性质，当任意个节点已经插入到BST后，再插入一个元素时，它的位置是唯一的
				所以不用关心左右子树是怎么排列的，而只用关心左、右节点分别有多少个
			5.多体会上一条的分析，可以写出任意情况的递推公式
				任意节点node的总数 = node的左子树的情况数 * node右子树的情况数
	3.可以添加记忆化数组来记录任意值 i 对应的BST总数：memo[i]，实现在递归时重复出现 n 值的情况
		边界条件：n = 1时，总数为 1
		补丁：当节点总数为 0 时，我们记为 1
dp：和括号生成的dp类似
思路：由递归的代码，很容易写出dp方程
	dp[0],dp[1] = 1,1
	for i := 0; i < n; i++ {	// 除去根节点
		dp[i] += dp[i] * dp[n-1+i]	// 左子树个数 * 右子树个数
	}
*/
//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	// 数学：lc，卡塔兰数
	//C := 1
	//for i := 0; i < n; i++ {
	//	C = C * 2 * (2*i + 1) / (i + 2)
	//}
	//return C

	// dp：根据递归，写出dp方程
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[i-j-1] * dp[j]
		}
	}
	return dp[n]

	// 个人方法
	//memo := make([]int, n+1)
	//memo[0], memo[1] = 1, 1
	//numTreesRecursion(n, memo)
	////fmt.Println(memo)
	//return memo[n]
	////return numTreesRecursion(n, memo)
}
var count = 1
func numTreesRecursion(n int, memo []int) int {
	if memo[n] == 0 {
		for i := 0; i < n; i++ {
			memo[n] += numTreesRecursion(n-1-i, memo) * numTreesRecursion(i, memo)
		}
		count++
	}
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)
