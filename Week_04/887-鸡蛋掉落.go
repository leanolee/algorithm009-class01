//你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N 共有 N 层楼的建筑。 
//
// 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。 
//
// 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。 
//
// 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。 
//
// 你的目标是确切地知道 F 的值是多少。 
//
// 无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？ 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入：K = 1, N = 2
//输出：2
//解释：
//鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
//否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
//如果它没碎，那么我们肯定知道 F = 2 。
//因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
// 
//
// 示例 2： 
//
// 输入：K = 2, N = 6
//输出：3
// 
//
// 示例 3： 
//
// 输入：K = 3, N = 14
//输出：4
// 
//
// 
//
// 提示： 
//
// 
// 1 <= K <= 100 
// 1 <= N <= 10000 
// 
// Related Topics 数学 二分查找 动态规划
package main

import (
	"../../../MyTools"
	"fmt"
)

func main() {
	K := 3
	N := 100
	drop := superEggDrop(K, N)
	fmt.Println(drop)
}

/*
dp：
数学法：
	参考：https://leetcode-cn.com/problems/super-egg-drop/solution/ji-dan-diao-luo-by-leetcode-solution-2/
	Go参考：https://leetcode-cn.com/problems/super-egg-drop/solution/gobian-ma-jie-fa-shuang-1000ms-19mb-by-ilwwli/
*/
//leetcode submit region begin(Prohibit modification and deletion)
func superEggDrop(K int, N int) int {
	// 组合？
	//N++
	//for base := 0; ; base++ { // 起扔楼层
	//	maxN := 0 // 最大到达楼层 maxheight
	//	for i := 0; i <= K && i <= base; i++ {
	//		maxN += comb(i, base)
	//	}
	//	if maxN >= N {
	//		return base
	//	}
	//}

	// 数学法：暂时放弃
	//dp := make([][]int, N+1)
	//for i, _ := range dp {
	//	dp[i] = make([]int, K+1)
	//	dp[i][1] = i
	//}
	//for i := 1; i <= K; i++ {
	//	dp[1][i] = 1
	//}
	//i := 2
	//for ; i <= N; i++ {
	//	for j := 2; j <= K; j++ {
	//		dp[i][j] = 1 + dp[i-1][j-1] + dp[i-1][j]
	//	}
	//	if dp[i][K] >= N {
	//		break
	//	}
	//}
	//return i

	// 决策单调性：上一种方法的演变
	dpLastX := make([]int, N+1)
	for i := 0; i <= N; i++ { // 一个鸡蛋时，需要扔鸡蛋次数即为最大楼层数
		dpLastX[i] = i
	}
	dpCurrX := make([]int, N+1) // 相当于二维数组：行：鸡蛋数；列：楼层数。维护一个 K*N 二维数组
	for i := 2; i <= K; i++ {   // i 个鸡蛋时
		//x := 1
		for x, n := 1, 1; n <= N; n++ { // x：起始扔鸡蛋的楼层假设为 1；n：总楼层数
			/*
				计算 i 个鸡蛋，总楼层 n 时，最佳开始扔鸡蛋的位置 x：
				比较当前 x 与上一层楼 x+1 之间的值，通过这样，逼近 最大值最小
				当 Max(dpLastX[x-1], dpCurrX[n-x]) <= MyTools.Max(dpLastX[x], dpCurrX[n-x-1]) 时，说明 x 最接近交叉点
			*/
			for x < n && MyTools.Max(dpLastX[x-1], dpCurrX[n-x]) > MyTools.Max(dpLastX[x], dpCurrX[n-x-1]) { // x 楼层扔鸡蛋
				x++ // 寻找最佳扔鸡蛋的起始位置
			}
			dpCurrX[n] = MyTools.Max(dpLastX[x-1], dpCurrX[n-x]) + 1 // 记得 +1
		}
		dpCurrX, dpLastX = dpLastX, dpCurrX // dpLastX：代表解法一中的 T(K-1,x-1)中的K-1；dpCurrX：代表解法一中的 T(K,N-x)中的K
	}
	return dpLastX[N]

	// 动态规划：二分
	//dpMemo := make(map[int]int)
	//return eggRecursion(K, N, dpMemo)
}

// 计算组合数C^m_n：n 个中选 m 个
func comb(m, n int) int {
	if m > n {
		m = n
	}
	if m > n/2 {
		m = n - m
	}
	ans := 1
	for i := n - m + 1; i <= n; i++ {
		ans *= i
	}
	for i := 2; i <= m; i++ {
		ans /= i
	}
	return ans
}

/*
题解：K个鸡蛋，N层楼，假设最高的扔鸡蛋不会碎的楼层是 F（在F层扔鸡蛋不会碎，在F+1层扔鸡蛋会碎）
	通过任意扔鸡蛋的方式可以去求得 F 的值，比如一个鸡蛋可以从 1 楼往上扔，但最多要扔 N 次（F=N时）才能确定 F 的值
	假设扔 T 次可以确定 F 值，求T的最小值
思路：
	随意选一个楼层 x 扔鸡蛋，以确定鸡蛋不会碎的最高楼层需要丢鸡蛋的最小次数 T。即在 x 层扔一个鸡蛋，存在两种结果：
		碎：T(K-1,x-1)=t1
			剩 K-1 个鸡蛋，x-1 楼层
			因为碎了，所以现在需要往下去找楼层 F
			即：从原来的 K个蛋，N层楼 中找F，变成了 K-1个蛋，x-1层楼 中找 F 的问题
			假设还要扔 t1 次，表达为 t1=T(K-1,x-1)
				但还需要加上在 x 层扔碎的这个鸡蛋，所以最后扔的次数为：t1 + 1
		不碎：T(K,N-x)=t2
			剩 K 个鸡蛋，N-x 楼层
			分析同上：最后扔的次数为：t2 + 1
		需要让 t1 t2 的值，尽量接近/相等（为什么两个值越接近，平均扔的次数最小呢？请参见下面的双蛋问题）
		即需要找到一个最佳楼层，使得 t1==t2 或是使 t1 t2 最接近的楼层
			最后通过取两者之间的最大值 max(t1,t2) = T
	那么怎么找到最佳扔鸡蛋的起始楼层 x 呢？
	二分法：
		需求：1-N楼层中，依次遍历，计算 t1 t2，找出 t1 t2 最接近时的 楼层值x
		显然，使用二分从 t1 = T(K-1 1) T(K-1 2) T(K-1 3)...T(K-1 N) 中找出与之最接近的 t2 = T(K N-2) T(K N-3) T(K N-4)...T(K N-(N+1))
			二分能最有效 O(log n) 的找到 x 的值
二分思路：
	t1=T(K-1,x-1)：x 的递增函数
	t2=T(K,N-x)：x 的递减函数
		要计算最接近的 t1 t2，也就是两个函数 相交点 为 t1==t2
	而两个函数的相交点 x 不出意外，都是浮点数
	则通过寻找该 浮点数 左右 最接近的整数（low high），再取 low high 计算 max(t1,t2) 后的最小值
双蛋问题：假设楼层100，蛋有2个
	二分法：log100 < 7，也就是需要7个蛋，才能保证使用二分法，但现在有两个蛋
	开方法：根号100 = 10，即从第10层开始扔，最坏的情况需要扔 1+9 - 9+10 次（鸡蛋最高的不碎的楼层分别是 1 100）
		如果是这种解法对应 leetcode-887，次数就是最坏的 19 次
		那么怎么让扔鸡蛋的最坏情况平均下呢？
	归纳法：假设第一次从 n 层开始扔，没碎，第2次从 n-1 层开始扔，没碎...以此类推，就变成了：
		n,n-1,n-2,n-3,...,1
		比如在 n 层没碎，在n-3层碎了，则一共最多扔：1+1+1+n-3=n次
		假如在第一次扔时就碎了，即在n层碎了，则一共最多扔：n次（一直从n扔到1）
		求得 n+n-1+n-2+...+1 = n(n+1)/2，由于总层数为 100
		所以 n(n+1)/2 >= 100，n = 14
*/
func eggRecursion(K int, N int, dpMemo map[int]int) int {
	target := N*100 + K
	if v, ok := dpMemo[target]; ok { // 已缓存的值
		return v
	}
	T := 0
	if K == 1 || N < 2 { // 边界条件：只有一个蛋，有一层/0层
		T = N
	} else {
		low, high := 1, N
		for low+1 < high { // 找到最佳扔鸡蛋的起始楼层
			mid := (low + high) >> 1               // 代表二分的楼层，化为两个子问题
			t1 := eggRecursion(K-1, mid-1, dpMemo) // T(K-1,x-1) 这个值未知，所以需要求这个值，可以通过递归化为 更小的楼层值的问题
			t2 := eggRecursion(K, N-mid, dpMemo)
			switch {
			case t1 < t2: // t1 小，说明扔鸡蛋的起始楼层低了
				low = mid
			case t1 > t2:
				high = mid
			case t1 == t2:
				low, high = mid, mid
			}
		}
		// 写法一
		//lower := MyTools.Max(eggRecursion(K-1, low-1, dpMemo), eggRecursion(K, N-low, dpMemo))
		//higher := MyTools.Max(eggRecursion(K-1, high-1, dpMemo), eggRecursion(K, N-high, dpMemo))
		//if lower < higher {	// 经过测试，lower higher其实是同一个值
		//	fmt.Println("not the same !!!")
		//}
		//T = MyTools.Min(lower, higher) + 1	// 计算找到 F 的最小值：取最大最小值
		// 写法二
		T = MyTools.Max(eggRecursion(K-1, low-1, dpMemo), eggRecursion(K, N-low, dpMemo)) + 1 // 记得 +1
	}
	dpMemo[target] = T
	return T
}

//leetcode submit region end(Prohibit modification and deletion)
