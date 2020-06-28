学习笔记

## 分治+回溯+递归+动态规划

> 1. 动态规划和递归或者分治没有根本上的区别（关键看有无最优的子结构）
> 2. 共性：找到重复子问题
> 3. 差异性：最优子结构、中途可以淘汰次优解

#### 递归

#### 分治

#### 回溯

#### 动态规划

1. Wiki定义：https://en.wikipedia.org/wiki/Dynamic_programming

2. “simplifying a complicated problem by breaking it down into simpler sub-problems“

   > in a recursive manner

3. Divide & Conquer + Optimal substructure：分治 + 最优子结构

4. 复杂DP

   > 维度：二维、三维
   >
   > 有取舍最优子结构

#### 感触

> 本质：寻找重复性 ——> 计算机指令集

1. 人肉递归低效、很累
2. 找到最近最简方法，将其拆解成可重复解决的问题
3. 数学归纳法思维（抵制人肉递归的诱惑）

#### 自顶向下 & 自底向上

1. 自顶向下：对于初学者和面试，可以习惯先递归、分治，然后记忆化搜索，再转化为自底向上的循环是没有问题的

   > 习惯于哪种思维方式，就用习惯的思维方式进行分治和记忆化，再转化成递推

2. 自底向上：对于熟练的选手或者追求比较深厚的DP功力，尤其是计算机竞赛的时候，只要开始写递归所有的竞赛型选手全部都是直接开始 for 循环，也就是全部都自底向上的写循环开始递推

   > 这也是为什么DP最好翻译成动态递推的原因

## 参考链接

- [递归代码模板](https://shimo.im/docs/EICAr9lRPUIPHxsH)
- [分治代码模板](https://shimo.im/docs/zvlDqLLMFvcAF79A)
- [动态规划定义](https://en.wikipedia.org/wiki/Dynamic_programming)

## 动态规划关键点

> 子问题、状态定义、DP方程

1. 最优子结构

2. 存储中间状态：状态定义，面试时最难

   > 和分治的区别：分治一般可把这个状态放在递归里；而dp一般要开数组来存

3. 递推公式（美其名曰：状态转移方程或者DP方程）：DP方程，竞赛时最难

## 动态规划小结

1. 打破自己的思维惯性，形成机器思维
2. 理解复杂逻辑的关键
3. 也是职业进阶的要点要领

## MIT algorithm course

1. B站搜索：mit动态规划

2. 如：https://www.bilibili.com/video/av53233912?from=search&seid=2847395688604491997

3. mit：5 easy steps to DP：（动态规划关键点，就是这5点的浓缩）

   > 1.2.3：分治思想，找重复性，4：递归+记忆化，再DP
   >
   > 1. define subproblems
   > 2. guess(part of solution)
   > 3. relate subproblem solutions：merge
   > 4. resurse & memoize or build DP fable bottom-up
   > 5. solve original problem