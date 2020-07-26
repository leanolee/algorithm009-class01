学习笔记

# 毕业总结

## 数据结构与算法

### 数据结构

#### 一维

1. 基础：数组 array（string）、链表 linked list
2. 高级：栈 stack、队列 queue、双端队列 deque、集合 set、映射 map（hash or map）、etc

#### 二维

> 从一维泛化而来

1. 基础：树 tree、图 graph
2. 高级：二叉搜索树 binary search tree（red-black tree、AVL）、堆 heap、并查集 disjoint set、字典树 Trie、etc

#### 特殊

> 主要用于工程中特定场景

1. 位运算 Bitwise、布隆过滤器 BloomFilter
2. LRU Cache

### 算法

1. if-else、switch ——> branch
2. for、while loop ——> iteration
3. 递归 Recursion（Divide & Conquer、Backtrace）
4. 搜索 Search：深度优先搜索 Depth first search、广度优先搜索 Breadth first search、A*、etc
5. 动态规划 Dynamic Programming
6. 二分查找 Binary Search
7. 贪心 Greedy
8. 数学 Math、几何 Geometry

## 时间、空间复杂度

#### 时间复杂度

1. O(1)：Constant Complexity 常数复杂度

2. O(log n)：Logarithmic Complexity 对数复杂度

3. O(n)：Linear Complexity 线性时间复杂度

4. O(n^2)：N square Complexity 平方

5. O(n^3)：N cubic Complexity 立方

6. O(2^n)：Exponential Growth 指数

7. O(n!)：Factorial 阶乘

#### 空间复杂度

1. 数组的长度
2. 递归的深度
3. 两者最大值

## 第一周重点

1. 数组、链表、跳表（跳表：升维、空间换时间）

   |  操作   | 数组 | 普通链表 | 跳表（有序） |
   | :-----: | :--: | :------: | :----------: |
   | prepend | O(1) |   O(1)   |     O(1)     |
   | append  | O(1) |   O(1)   |     O(1)     |
   | lookup  | O(1) |   O(n)   |   O(log n)   |
   | insert  | O(n) |   O(1)   |   O(log n)   |
   | delete  | O(n) |   O(1)   |   O(log n)   |

2. 栈、队列、Deque、PriorityQueue

   1. Stack：先入后出；添加、删除皆为 O(1)，查询 O(n)
   2. Queue：先入先出；添加、删除皆为 O(1)，查询 O(n)
   3. Deque：两端可以进出的QueueDeque - Double-End ，插入和删除都是 O(1)，查询 O(n)
   4. PriorityQueue：插入操作：O(1)，取出操作：O(logN) - 按照元素的优先级取出，底层具体实现的数据结构较为多样和复杂：heap（二叉树堆、Fibonacci堆...）、bst（平衡二叉搜索树：红黑树、AVL...）、treap

## 第二周要点

1. 哈希表、映射、集合（底层：哈希表、二叉树-红黑树）

   | Data Structure |  Access   |  Search   | Insertion | Deletion  |   Worst   |   Space   |
   | :------------: | :-------: | :-------: | :-------: | :-------: | :-------: | :-------: |
   |   Hash Table   |    N/A    |   O(1)    |   O(1)    |   O(1)    |   O(n)    |   O(n)    |
   | Red-Black Tree | O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) |

2. 树、二叉树、二叉搜索树

   1. 遍历及代码模板：递归、栈、莫里斯

      > 前序（Pre-order）：根-左-右
      >
      > 中序（In-order）：左-根-右
      >
      > 后序（Post-order）：左-右-根

   2. BST

      > 查询：O(log n)
      >
      > 插入新节点（创建）：O(log n)
      >
      > 删除：O(log n)

3. 堆、二叉堆、图

   1. 操作

      >find-max：O(1)
      >
      >delete-max：O(logN)
      >
      >insert(create)：O(logN) or O(1)

   2. 实现

      > 索引为 i 的左孩子的索引是 (2**i + 1)
      >
      > 索引为 i 的右孩子的索引是 (2*i + 2)
      >
      > 索引为 i 的父节点的索引是 floor((i -1)/2)：(i - 1)>>1

   3. 延伸：Fibonacci堆

## 第三周要点

1. 默写代码模板

2. 递归：思维要点

   > 抵制人肉递归
   >
   > 找最近重复性
   >
   > 数学归纳法思维

3. 分治、回溯

   1. 最近重复性：分治、回溯，本质上就是种递归
   2. 最优重复性：动态规划

## 第四周要点

1. DFS、BFS：默写代码模板

2. 贪心：适用场景

   > 问题能够分解成子问题来解决，子问题的最优解能递推到最终问题的最优解
   >
   > 贪心：对每个子问题的解决方案都做出选择，不能回退
   >
   > 动态规划：会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能

3. 二分查找：默写代码模板

   1.  关键点：

      > l<r or l<=r
      >
      > l=mid or l=mid+1
      >
      > r=mid or r=mid-1

   2. 搜索一个元素：

      > 搜索区间两端闭
      > <=
      > if 相等返回
      > mid + 1 & mid - 1

   3. 搜索左右边界：

      > 左闭右开常见：例 l,r:=0,len(nums)，就是左闭右开
      > <
      > if 相等不返回
      > mid +- 与否，看区间开/闭
      > 判断返回

## 第六周要点

1. 分治、回溯、递归、动态规划

   1. 动态规划和递归或者分治没有根本上的区别（关键看有无最优的子结构）
   2. 共性：找到重复子问题
   3. 差异性：最优子结构、中途可以淘汰次优解

2. 关键点

   1. 最优子结构

   2. 存储中间状态：状态定义，面试时最难

      > 和分治的区别：分治一般可把这个状态放在递归里；而dp一般要开数组来存

   3. 递推公式（美其名曰：状态转移方程或者DP方程）：DP方程，竞赛时最难

3. 总结

   1. 打破自己的思维惯性，形成机器思维
   2. 理解复杂逻辑的关键
   3. 也是职业进阶的要点要领

## 第七周要点

1. Trie：默写代码模板

   1. 结点本身不存完整单词
   2. 从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串
   3. 每个节点的所有子节点路径代表的字符都不相同

2. 并查集：默写代码模板

   1. makeSet(s)

      > 建立一个新的并查集，其中包含 s 个单元素集合

   2. unionSet(x, y)

      > 把元素 x 和元素 y 所在的集合合并，要求 x 和 y 所在的集合不相交，如果相交则不合并

   3. find(x)

      > 找到元素 x 所在的集合的代表，该操作也可以用于判断两个元素是否位于同一个集合，只要将它们各自的代表比较一下就可以了

3. 高级搜索

   1. 优化：

      1. 朴素搜索

      2. 优化方式：不重复（fibonacci）、剪枝（生成括号问题）

      3. 搜索方向：

         > DFS：depth first search 深度优先搜索
         >
         > BFS：breadth first search 广度有限搜索
         >
         > 双向搜索、启发式搜索

   2. Two-ended BFS：默写代码模板

   3. A*：估价函数

      1. 启发式函数： h(n)，它用来评价哪些结点最有希望的是一个我们要找的结点，h(n) 会返回一个非负实数，也可以认为是从结点n的目标结点路径的估计成本
      2. 启发式函数是一种告知搜索方向的方法。它提供了一种明智的方法来猜测哪个邻居结点会导向一个目标

4. 红黑树、AVL

   1. AVL

      1. 平衡二叉搜索树

      2. 每个结点存 balance factor = { 1, 0, 1}

      3. 四种旋转操作

      4. 不足：结点需要存储额外信息、且调整次数频繁

         > 其他的近似平衡二叉搜索树

   2. 红黑树

      1. 每个结点要么是红色，要么是黑色
      2. 根结点是黑色
      3. 每个叶结点（ NIL结点，空结点）是黑色的 
      4. 不能有相邻接的两个红色结点
      5. 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点

   3. 操作要点：

      1. 左旋
      2. 右旋
      3. 左右旋
      4. 右左旋

   4. 使用场景

      1. 使用AVL：读操作非常多，写操作很少

         > DataBase

      2. 使用红黑树：更多插入、删除

         > 高级语言的库，map、set

   5. 对比：

      - AVL trees provide faster lookups than Red Black Trees because they are more strictly
        balanced.
      - Red Black Trees provide faster insertion and removal operations than AVL trees as
        fewer rotations are done due to relatively relaxed balancing.
      - AVL trees store balance factors or heights with each node, thus requires storage for
        an integer per node whereas Red Black Tree requires only 1 bit of information per
        node.
      - Red Black Trees are used in most of the language libraries
        like map, multimap, multisetin C++ whereas AVL trees are used in databases where
        faster retrievals are required.

## 第八周要点

#### 位运算

1. 位运算：<<，>>，|，&，~，^

2. 位置运算

   1. 将 x 最右边的 n 位清零： x & (~0 << n）
   2. 获取 x 的第 n 位值（0 或者 1）： (x >> n) & 1
   3. 获取 x 的第 n 位的幂值： x & (1 << n）
   4. 仅将第 n 位置为 1： x | (1 << n)
   5. 仅将第 n 位置为 0： x & (~ (1 << n))
   6. 将 x 最高位至第 n 位（含）清零： x & ((1 << n) - 1)
   7. 将第 n 位至第 0 位（含）清零： x & (~ ((1 << (n + 1)) - 1))

3. 操作技巧

   1. 判断奇偶

      > x % 2 == 1 ——> (x & 1) == 1
      > x % 2 == 0 ——> (x & 1) == 0

   2. x >> 1 ——> x / 2

      > 即： x = x / 2; ——> x = x >> 1;
      > mid = (left + right) / 2; ——> mid = (left + right) >> 1;

   3. X = X & (X - 1)： 清零最低位的 1

   4. X & -X： 得到最低位的 1

   5. X & ~X： 0

4. 取最高位：

   ```go
   c := a ^ b
   c = c | (c >> 1)
   c = c | (c >> 2)
   c = c | (c >> 4)
   c = c | (c >> 8)
   c = c | (c >> 16)
   c = c>>1 + 1
   ```

#### 布隆过滤器 & LRU Cache

1. 布隆过滤器

   1. 一个很长的二进制向量和一系列随机映射函数。布隆过滤器可以用于检索一个元素是否在一个集合中

   2. 优点：是空间效率和查询时间都远远超过一般的算法

   3. 缺点：是有一定的误识别率和删除困难

      > 快速查询的缓存：
      >
      > 不存在：100%
      >
      > 存在：可能存在

2. LRU Cache：哈希表+双向链表

   1. 两个要素： 大小 、替换策略
   2. Hash Table + Double LinkedList
   3. O(1) 查询，O(1) 修改、更新

#### 排序算法

1. 时间 & 空间复杂度

   | 排序方法 | 时间复杂度-平均 | 时间复杂度-最坏 | 时间复杂度-最好 |  空间复杂度  | 稳定性 |
   | :------: | :-------------: | :-------------: | :-------------: | :----------: | :----: |
   | 插入排序 |   O($ n^2 $)    |   O($ n^2 $)    |      O(n)       |     O(1)     |  稳定  |
   | 希尔排序 | O($ n^{1.3} $)  |   O($ n^2 $)    |      O(n)       |     O(1)     | 不稳定 |
   | 选择排序 |   O($ n^2 $)    |   O($ n^2 $)    |   O($ n^2 $)    |     O(1)     | 不稳定 |
   |  堆排序  |  O($nlog_2n$)   |  O($nlog_2n$)   |  O($nlog_2n$)   |     O(1)     | 不稳定 |
   | 冒泡排序 |   O($ n^2 $)    |   O($ n^2 $)    |      O(n)       |     O(1)     |  稳定  |
   | 快速排序 |  O($nlog_2n$)   |   O($ n^2 $)    |  O($nlog_2n$)   | O($nlog_2n$) | 不稳定 |
   | 归并排序 |  O($nlog_2n$)   |  O($nlog_2n$)   |  O($nlog_2n$)   |     O(n)     |  稳定  |
   |          |                 |                 |                 |              |        |
   | 计数排序 |     O(n+k)      |     O(n+k)      |     O(n+k)      |    O(n+k)    |  稳定  |
   |  桶排序  |     O(n+k)      |   O($ n^2 $)    |      O(n)       |    O(n+k)    |  稳定  |
   | 基数排序 |     O(n*k)      |     O(n*k)      |     O(n*k)      |    O(n+k)    |  稳定  |

2. 详见：

   初级排序算法

   [https://github.com/leanolee/algorithm009-class01/blob/master/Week_08/4.%E5%88%9D%E7%BA%A7%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95.md](https://github.com/leanolee/algorithm009-class01/blob/master/Week_08/4.初级排序算法.md)

   高级排序算法

   [https://github.com/leanolee/algorithm009-class01/blob/master/Week_08/5.%E9%AB%98%E7%BA%A7%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95.md](https://github.com/leanolee/algorithm009-class01/blob/master/Week_08/5.高级排序算法.md)

   特殊排序算法

   [https://github.com/leanolee/algorithm009-class01/blob/master/Week_08/6.%E7%89%B9%E6%AE%8A%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95.md](https://github.com/leanolee/algorithm009-class01/blob/master/Week_08/6.特殊排序算法.md)

## 第九周要点

#### 高级动态规划

1. “Simplifying a complicated problem by breaking it down into simpler sub-problems”
   (in a recursive manner - 分治思想)

2. Divide & Conquer + Optimal substructure：分治 + 最优子结构
3. 顺推形式： 动态递推

#### 字符串算法

> 详细参考：[https://github.com/leanolee/algorithm009-class01/blob/master/Week_010/%E5%AD%97%E7%AC%A6%E4%B8%B2%E5%8C%B9%E9%85%8D%E7%AE%97%E6%B3%95%E5%90%88%E9%9B%86.md](https://github.com/leanolee/algorithm009-class01/blob/master/Week_010/字符串匹配算法合集.md)

1. bruteForce：蛮力匹配

2. Sunday算法

   1. 前缀字符串匹配

   2. 位移表（坏字符）规则：后移距离 = 最右边不重复字符到字符串末尾的距离

      > 不存在的字符：全部移动 len(p)+1 位
      >
      > 存在的字符：移动位数 = len(p) - 字符最后一次出现的索引

3. KMP算法

   1. KMP算法（Knuth-Morris-Pratt）的思想就是，当子串与目标字符串不匹配时，其实你已经知道了前面已经匹配成功那一部分的字符（包括子串与目标字符串）

      > 以阮一峰的文章为例，当空格与 D 不匹配时，你其实 知道前面六个字符是“ABCDAB”

   2. KMP 算法的想法是，设法利用这个已知信息，不要把“搜索位置” 移回已经比较过的位置，继续把它向后移，这样就提高了效率

      > 最长重复的前缀和后缀

4. Boyer-Moore算法

   1. 坏字符规则：
      1. 不存在的字符
      2. 存在的字符
   2. 好后缀规则：
      1. 存在与好后缀完全匹配的子串
      2. 存在与好后缀部分匹配的子串（前缀）
      3. 不存在与好后缀匹配的子串

5. Rabin-Karp算法

   1. 通过哈希函数，我们可以算出子串的哈希值，然后将它和目标字符串中的子串的哈希值进行比较。 这个新方法在速度上比暴力法有显著提升

   2. API的hash函数，是O(n)的
   3. 比较 hash 值：

     > 如果 hash 值不同，字符串必然不匹配
     >
     > 如果 hash 值相同，还需要使用朴素算法再次判断

