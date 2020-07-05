//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。 
//
// 示例 1: 
//
// 
//输入:
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出: 3
//解释: 
//长度最长的公共子数组是 [3, 2, 1]。
// 
//
// 说明: 
//
// 
// 1 <= len(A), len(B) <= 1000 
// 0 <= A[i], B[i] < 100 
// 
// Related Topics 数组 哈希表 二分查找 动态规划
package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}
	//A := []int{0, 0, 0, 0, 0}
	//B := []int{0, 0, 0, 0, 0}
	length := findLength(A, B)
	fmt.Println(length)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findLength(A []int, B []int) int {
	// 二分 + hash：Rabin-Karp 算法
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	base, mod := 101, 1000000009 // base（根据题意，选择大于 100 的质数） 和 mod 的选择
	pow := func(x, n int) int { // 用于计算 mult，固定长度的 hash 乘因子
		ans := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				//ans *= x
				ans = ans * x % mod
			}
			//x *= x
			x = x * x % mod
		}
		return ans
	}
	check := func(mid int) bool {
		hashA, memoA, mult := 0, make(map[int]bool), pow(base, mid-1)
		for i := 0; i < len(A); i++ {
			if i < mid {
				hashA = (hashA*base + A[i]) % mod // 累积 hash 值
			} else { // 滑动窗口 hash 值
				memoA[hashA] = true // 上一次的 hashA
				//hashA = (hashA-A[i-mid]*mult)*base + A[i]	// 不 mod 时
				hashA = ((hashA-A[i-mid]*mult%mod+mod)%mod*base + A[i]) % mod
			}
		}
		memoA[hashA] = true

		hashB := 0
		for i := 0; i < len(B); i++ {
			if i < mid {
				hashB = (hashB*base + B[i]) % mod
			} else {
				if memoA[hashB] { // 查看是否有 重复子数组
					return true
				}
				hashB = ((hashB-B[i-mid]*mult%mod+mod)%mod*base + B[i]) % mod
			}
		}
		if memoA[hashB] {
			return true
		}
		return false
	}
	// lc
	l, r := 1, min(len(A), len(B))+1
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// 个人写法
	//l, r := 0, min(len(A), len(B))
	//for l <= r {
	//	mid := (l + r) >> 1
	//	if check(mid) {
	//		l = mid + 1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	return l - 1

	// DP：F(i,j) = F(i-1,j-1)+1 / F(i,j) = 0
	//max, m, n := 0, len(A), len(B)
	//dp := make([][]int, m+1)
	//dp[0] = make([]int, n+1)
	//for i := 1; i <= m; i++ {
	//	dp[i] = make([]int, n+1)
	//	for j := 1; j <= n; j++ {
	//		if A[i-1] == B[j-1] {	// F(i,j) = F(i-1,j-1)+1
	//			dp[i][j] = dp[i-1][j-1] + 1
	//			if dp[i][j] > max {
	//				max = dp[i][j]
	//			}
	//		}	// else：F(i,j) = 0
	//	}
	//}
	//return max

	// 滑动窗口： O((N + M)*min(N,M))
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//max, m, n := 0, len(A), len(B)
	//getMax := func(aIdx, bIdx int, length int) int {
	//	curr, k := 0, 0 // 最长重复子数组、当前最长
	//	for i := 0; i < length; i++ {
	//		if A[aIdx+i] == B[bIdx+i] {
	//			k++
	//			curr = maxVal(curr, k)
	//		} else {
	//			k = 0
	//		}
	//	}
	//	return curr
	//}
	//for i := 0; i < m; i++ { // 固定 B，滑动 A
	//	length := minVal(m-i, n)
	//	if length > max { // 剪枝：如果最短的没有 max 大，就不用找了
	//		max = maxVal(max, getMax(i, 0, length))
	//	}
	//}
	//for i := 0; i < n; i++ { // 固定 A，滑动 B
	//	length := minVal(n-i, m)
	//	if length > max {
	//		max = maxVal(max, getMax(0, i, length))
	//	}
	//}
	//return max

	// 暴力
}

//leetcode submit region end(Prohibit modification and deletion)
