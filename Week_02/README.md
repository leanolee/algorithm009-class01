学习笔记

# 周总结-week02

## 本周目标

1. 哈希表、映射、集合
2. 树、二叉树、二叉搜索树
3. 堆和二叉堆、图

## 重点题型

#### 49-字母异位词分组

1. 通过该题，体会hash的用法

2. 思路：

   ```go
   /*
   hash：
      注意：write 字符的顺序不同，则最后的hash值不同
      算法：
         将字符串中字符出现的频次，放入一个 26长度[]int 中，再依次取出做hash
         用 hash值 作为 key，将异位词都放进该分组下
   */
   func groupAnagrams(strs []string) [][]string {
       memo := [26]int{}
   	group := make(map[uint32][]string)
   	new32 := fnv.New32()
   	for _, str := range strs {
   		for _, b := range str {
   			memo[b-'a'] = memo[b-'a'] + 1
   		}
   		hashByte := make([]byte, 0)
   		for i, count := range memo {
   			if count == 0 {
   				continue
   			}
   			for j := 0; j < count; j++ {
   				hashByte = append(hashByte, byte('a'+i))
   			}
   			memo[i] = 0
   		}
   		new32.Write(hashByte)
   		sum32 := new32.Sum32()
   		new32.Reset()
   		group[sum32] = append(group[sum32], str)
   	}
   	anagrams := make([][]string, len(group))
   	i := 0
   	for _, v := range group {
   		anagrams[i] = v
   		i++
   	}
   	return anagrams
   }
   ```

#### 二叉树/N叉树的遍历

> 参见总结：
>
> [https://github.com/leanolee/algorithm009-class01/blob/master/Week_02/%E6%A0%91%E7%9A%84%E9%81%8D%E5%8E%86.md](https://github.com/leanolee/algorithm009-class01/blob/master/Week_02/树的遍历.md)

#### 347-前 K 个高频元素

1. 本题可使用方法：

   > 暴力法：O(nk),O(n)
   >
   > API+数组
   >
   > API+hash
   >
   > 小顶堆：O(nlogk),O(n)
   >
   > 快排：减治法
   >
   > 1. 参见：https://zhuanlan.zhihu.com/p/76734219
   > 2. 体会面试时非算法的软知识

2. 通过该题的学习

   1. 手写大、小顶堆：重要的 insert，delete 方法
   2. 学习了Go源码的二叉堆（待补充心得）
   3. 手写快排/随机快排实现：重要 partition 方法
   4. 学习Go源码的排序：Sort，stable（待完善学习）

#### 51-N皇后

1. 结合leetcode-37：解数独，学习回溯（以及位运算）

2. 对比递归时，使用回溯和不适用回溯的情况

   ```go
   对比：leetcode-51 leetcode-37
      存数据：
         37：已使用：0，未使用：1
         51：已使用：0，未使用：1
      使用数据：
         37：已使用：1，未使用：0
         51：已使用：0，未使用：1
      遍历：
         37：轮询 num，去 位存的数据 中查
         51：轮询 位数据，获取能使用的 num
      回溯：
         37：使用 数组 作为递归参数，需要回溯。因为改变 int 后，还要放进数组中
         51：使用 int 作为递归参数，不需要回溯
   ```

## 本周不足

1. 重要：未完成 hashmap 源码的研读
2. 非常重要：leetcode 优秀题解未学习很多，尤其国际站
3. 重要：部分题没有进行第3次，甚至第2次的练习