//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。 
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ 
//
// 注意：给定 n 是一个正整数。 
//
// 示例 1： 
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶 
//
// 示例 2： 
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
// 
// Related Topics 动态规划
package main

import "fmt"

func main() {
	n := 9
	//stairs := climbStairs(n)
	//stairs := noRepeatDP(n)
	stairs := climbStairsNoRepeat(n)
	fmt.Println(stairs)
}

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 2
	return climbStairsRecursion(n-1, memo)
}

func climbStairsRecursion(n int, memo []int) int {
	if memo[n] == 0 {
		memo[n] = climbStairsRecursion(n-1, memo) + climbStairsRecursion(n-2, memo)
	}
	return memo[n]
}
func noRepeatDP(n int) int {
	// 使用升维思想再做一遍
	steps := []map[int]int{{1: 1, 2: 0, 3: 0}, {1: 0, 2: 1, 3: 0}, {1: 1, 2: 1, 3: 1}}
	if n < 4 {
		if n == 3 {
			return n
		}
		return 1
	}
	for i := 4; i <= n; i++ {
		steps[0], steps[1], steps[2] = steps[1], steps[2], steps[0]
		steps[2][1], steps[2][2], steps[2][3] = steps[1][2]+steps[1][3], steps[0][1]+steps[0][3], steps[2][1]+steps[2][2]
	}
	return steps[2][1] + steps[2][2] + steps[2][3]
}
func climbStairsNoRepeat(n int) int {
	memo := make([][3]int, n+1)
	//memo[0], memo[1], memo[2] = [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{1, 1, 1}
	memo[0], memo[1], memo[2], memo[3] = [3]int{0, 0, 0}, [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{1, 1, 1}
	return climbStairsNoRepeatRecursion(n-1, memo, 0) + climbStairsNoRepeatRecursion(n-2, memo, 1) + climbStairsNoRepeatRecursion(n-3, memo, 2)
}
func climbStairsNoRepeatRecursion(n int, memo [][3]int, end int) int {
	noRepeat := 0
	for i := 0; i <= 2; i++ {
		if i != end && n > i {
			if memo[n-i-1][i] == 0 {
				memo[n-i-1][i] = climbStairsNoRepeatRecursion(n-i-1, memo, i)
			}
			noRepeat += memo[n-i-1][i]
		}
	}
	return noRepeat
}

//func dp(n int, step int) int {
//	if n < 0 {
//		return 0
//	}
//	if n == 0 {
//		return 1
//	}
//	ans := 0
//	for i := 1; i <= 3; i++ {
//		if i != step {
//			ans += dp(i-step, step)
//		}
//	}
//	return ans
//}
//func climbStairsNoRepeat_(n int) int {
//	return dp(n-1, 1) + dp(n-2, 2) + dp(n-3, 3)
//}

//leetcode submit region end(Prohibit modification and deletion)
/*
######################
ruizhe
# 3步爬楼梯不重复.py


# 思考
# 如果没有不能同步数限制
# f(n) = f(n-3) + f(n-2) + f(n-1)

# 但是需要记录额外信息，即上一次走了几步

# f(n) = f(n-3) if last_step != 3
#      + f(n-2) if last_step != 2
#      + f(n-1) if last_step != 1

# 所以需要记录两个状态
# f(i, j) i表示走到第i个台阶，j表示这一步走了j个台阶

# f(i, 1) = f(i-1, 2) + f(i-1, 3)
# f(i, 2) = f(i-2, 1) + f(i-2, 3)
# f(i, 3) = f(i-3, 1) + f(i-3, 2)

# f(i)= sum f(i, j) j = 1, 2, 3

# f(1, 1) = 1
# 1, 2 = 0
# 1, 3 = 0
# 2, 1 = 0
# 2, 2 = 1
# 2, 3 = 0
# 3, 1 = 1
# 3, 2 = 1
# 3, 3 = 1


class Solution:
memo = {
(1, 1): 1,
(1, 2): 0,
(1, 3): 0,
(2, 1): 0,
(2, 2): 1,
(2, 3): 0,
(3, 1): 1,
(3, 2): 1,
(3, 3): 1
}

def climbStairs(self, n: int) -> int:
if n <= 0:
return 0
return self.helper(n, 1) + self.helper(n, 2) + self.helper(n, 3)

def helper(self, n, k) -> int:

if (n, k) in self.memo:
return self.memo[(n, k)]

if k == 1:
self.memo[(n, k)] = self.helper(n - k, 2) + self.helper(n - k, 3)
elif k == 2:
self.memo[(n, k)] = self.helper(n - k, 1) + self.helper(n - k, 3)
else:
self.memo[(n, k)] = self.helper(n - k, 1) + self.helper(n - k, 2)

return self.memo[(n, k)]


def main():
s = Solution()
res = s.climbStairs(1)
print(res)


if __name__ == '__main__':
main()




######################

class Solution:
def climbStairs(self, n: int) -> int:
@lru_cache()
def dp(ii, invalid_step):
if ii < 0:
return 0
if ii == 0:
return 1

ans = 0
for step in [1, 2, 3]:
if step != invalid_step:
ans += dp(ii-step, step)
return ans

return dp(n-1, 1) + dp(n-2, 2) + dp(n-3, 3)


int first = 1;  // 第一个值
int second = 1; // 第二个值
int third = 0;  // 第三个值
// 包括第n个
for( int i = 1 ; i <= n ; i++ ) {
if ( i == 1 || i == 2) {
third = 1；
} else {
// 循环覆盖前面两个值
third  = first + second ;
first  = second;
second = third;
}
}
System.out.println(second);

*/
