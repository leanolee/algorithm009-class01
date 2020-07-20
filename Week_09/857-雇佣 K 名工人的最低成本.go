//有 N 名工人。 第 i 名工人的工作质量为 quality[i] ，其最低期望工资为 wage[i] 。 
//
// 现在我们想雇佣 K 名工人组成一个工资组。在雇佣 一组 K 名工人时，我们必须按照下述规则向他们支付工资： 
//
// 
// 对工资组中的每名工人，应当按其工作质量与同组其他工人的工作质量的比例来支付工资。 
// 工资组中的每名工人至少应当得到他们的最低期望工资。 
// 
//
// 返回组成一个满足上述条件的工资组至少需要多少钱。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入： quality = [10,20,5], wage = [70,50,30], K = 2
//输出： 105.00000
//解释： 我们向 0 号工人支付 70，向 2 号工人支付 35。 
//
// 示例 2： 
//
// 输入： quality = [3,1,10,10,1], wage = [4,8,2,2,7], K = 3
//输出： 30.66667
//解释： 我们向 0 号工人支付 4，向 2 号和 3 号分别支付 13.33333。 
//
// 
//
// 提示： 
//
// 
// 1 <= K <= N <= 10000，其中 N = quality.length = wage.length 
// 1 <= quality[i] <= 10000 
// 1 <= wage[i] <= 10000 
// 与正确答案误差在 10^-5 之内的答案将被视为正确的。 
// 
// Related Topics 堆
package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	//quality := []int{3, 1, 10, 10, 1}
	//wage := []int{4, 8, 2, 2, 7}
	//k := 3
	//quality := []int{10, 20, 5}
	//wage := []int{70, 50, 30}
	//k := 2
	quality := []int{25, 68, 35, 62, 52, 57, 35, 83, 40, 51}
	wage := []int{147, 97, 251, 129, 438, 443, 120, 366, 362, 343}
	k := 6 // 答案：1979.31429 vs 2086.8857142857146
	workers := mincostToHireWorkers(quality, wage, k)
	fmt.Println(workers)
}
/*
原生heap的使用：详见笔记：Go原生Heap使用方法
*/
//leetcode submit region begin(Prohibit modification and deletion)
func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	// 堆O(n long n)
	n := len(quality)
	workers := Workers(make([]Worker, n))
	for i := 0; i < n; i++ {
		workers[i] = Worker{quality[i], wage[i], float64(wage[i]) / float64(quality[i])}
	}
	sort.Sort(&workers)
	ans := math.MaxFloat64
	h, sumQ := &QHeap{}, 0
	//h, sumQ := QHeap(quality), 0
	for _, w := range workers {
		h.Push(-w.quality)
		heap.Fix(h, h.Len()-1)
		sumQ += w.quality
		if h.Len() > K {
			//sumQ += h.Pop().(int)	// 写法二
			//heap.Fix(h, 0)
			sumQ += (*h)[0] // 写法三
			heap.Remove(h, 0)
		}
		//heap.Init(h)	// 写法一
		if h.Len() == K {
			ans = getMin(ans, float64(sumQ)*(w.ratio))
		}
	}
	return ans

	// 暴力 + 贪心：O(n^2 long n)，超时
	N := len(quality)
	cost, work, maxCost := make([]float64, N), make([]float64, 0), math.MaxFloat32
	for i := 0; i < N; i++ { // 计算性价比
		cost[i] = float64(wage[i]) / float64(quality[i])
	}
	fmt.Println(cost)
	for _, c := range cost {
		for i, cw := range cost { // 获取满足当前 性价比 的人员
			if cw <= c {
				work = append(work, c*float64(quality[i]))
			}
		}
		if len(work) >= K {
			sort.Float64s(work) // 必须排序
			var total float64 = 0
			for i := 0; i < K; i++ { // 计算 总cost
				total += work[i]
			}
			if total < maxCost {
				maxCost = total
			}
		}
		work = work[:0]
	}
	return maxCost
}

type Workers []Worker

func (w *Workers) Len() int {
	return len(*w)
}
func (w *Workers) Less(i, j int) bool {
	return (*w)[i].ratio < (*w)[j].ratio
}
func (w *Workers) Swap(i, j int) {
	(*w)[i], (*w)[j] = (*w)[j], (*w)[i]
}

type QHeap []int

func (q *QHeap) Len() int {
	return len(*q)
}
func (q *QHeap) Less(i, j int) bool {
	return (*q)[i] < (*q)[j]
}
func (q *QHeap) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}
func (q *QHeap) Pop() (v interface{}) {
	//*q, v = (*q)[:q.Len()-1], (*q)[q.Len()-1]
	*q, v = (*q)[1:], (*q)[0]
	return
}
func (q *QHeap) Push(v interface{}) {
	*q = append(*q, v.(int))
}

type Worker struct {
	quality int
	wage    int
	ratio   float64
}

func getMin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
