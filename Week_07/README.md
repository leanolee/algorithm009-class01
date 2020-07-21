学习笔记

## 作业1

#### 目标


> 分析单词搜索 2 用 Tire 树方式实现的时间复杂度，请同学们提交在学习总结中

#### 分析

```go
/* 
   N:words数组中字符串个数
   k:字符串的最大长度
   m:矩阵的行数
   n:矩阵的列数
   L:单词最大的长度
时间复杂度：
   创建Trie：O(N*k)
   遍历矩阵：分两种情况讨论
情况一：m,n 为任意值，单词的长度任意长，字符串数组中字符串个数任意多
   O(m*n * 4^(m*n)) 或 O(m*n * 4*3^(m*n-1))
      1.遍历矩阵：m*n
      2.DFS：由于字符串数组中个数任意多，所以可以认为矩阵中的每个元素都会被遍历，每次4个方向
         4^(m*n)
      3.由于每次扩展4个方向时，有一个是起始方向可以忽略（但是第一次扩展是4个方向）
         4*3^(m*n-1)
情况二：单词长度最大为L
   O(m*n * 4^L) 或 O(m*n * 4*3^(L-1))
      1.遍历矩阵：m*n
      2.DFS：由于单词长度最大为L，所以当遍历到L个元素后，就可以终止
         4^L
      3.由于每次扩展4个方向时，有一个是起始方向可以忽略（但是第一次扩展是4个方向）
         4*3^(m*n-1)
*/
```

## 作业2

#### 目标

> 总结双向 BFS 代码模版

#### 模板

```go
// Go
func twoEndedBFS(graph, head, tail) {
    begin, end, beginVisited, endVisited := slice[]{head}, slice[]{tail}, map[]{head}, map[]{tail}
    for len(begin) > 0 {
        temp := slice[]{head}
        for k, node := range begin {
            if beginVisited[k] {
                continue
            }
            beginVisited[k]
            process(k)
            nodes = generate_related_nodes(node)
            temp[nodes...]
        }
        begin = temp
        if len(begin) > len(end) {
            begin, end, beginVisited, endVisited = end, begin, endVisited, beginVisited
    	}
    }
}
```

