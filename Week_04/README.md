学习笔记

# DFS、BFS、贪心算法、二分查找

### 求旋转数组的旋转起始位置

#### 需求

> 一个旋转数组，找到它旋转的起始位置

#### 思路

1. 迭代：O(n)，当有重复元素时，只能使用迭代

   ```go
   func idxOfHalfSort(nums []int) int {
   	for i := 1; i < len(nums); i++ {
   		if nums[i] < nums[i-1] {
   			return nums[i]
   		}
   	}
   	return nums[0]
   }
   ```

2. 二分：旋转数组寻找最小值，使用二分的前提条件是，数组不包含重复元素

   > 三种判断可以确定：
   >
   > 1. 判断当前值大于它右边的值，返回右边的值
   > 2. 判断当前值小于它左边的值，返回当前值
   > 3. 判断首小于尾，返回首
   > 4. 左侧逼近最大值，右侧逼近最小值，最后返回左侧+1

#### 方法一

1. 判断当前值大于它右边的值，返回右边的值

2. 代码

   ```go
   func idxOfHalfSort(nums []int) int {
   	l, r := 0, len(nums)
   	for l < r {
   		mid := (l + r) >> 1
   		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
   			return nums[mid+1]
   		}
   		if nums[mid] > nums[l] {
   			l = mid
   		} else {
   			r = mid - 1
   		}
   	}
   	return nums[r+1]
   }
   ```

#### 方法二

1. 判断当前值小于它左边的值，返回当前值

2. 代码

   ```go
   func idxOfHalfSort(nums []int) int {
   	l, r := 0, len(nums)
   	for l < r {
   		mid := (l + r) >> 1
   		if mid > 0 && nums[mid] < nums[mid-1] {
   			return nums[mid]
   		}
   		if nums[mid] > nums[l] {
   			l = mid
   		} else {
   			r = mid
   		}
   	}
   	return nums[l]
   }
   ```

#### 方法三

1. 判断首小于尾，返回首

2. 代码

   ```go
   func idxOfHalfSort(nums []int) int {
   	l, r := 0, len(nums)-1
   	for l < r {
   		mid := (l + r) >> 1
   		if nums[l] < nums[r] {
   			return nums[l]
   		}
   		if nums[mid] >= nums[l] {
   			l = mid + 1
   		} else {
   			r = mid
   		}
   	}
   	return nums[r]
   }
   ```

#### 方法四

1. 左侧逼近最大值，右侧逼近最小值，最后返回左侧+1

2. 代码

   ```go
   func idxOfHalfSort(nums []int) int {
   	l, r := 0, len(nums)
   	if r == 1 {
   		return nums[0]
   	}
   	for l < r {
   		mid := (l + r) >> 1
   		if nums[mid] > nums[l] {
   			l = mid
   		} else {
   			r = mid
   		}
   	}
   	return nums[l+1]
   }
   ```

### DFS

### BFS

### 贪心算法

### 二分算法

待补充...