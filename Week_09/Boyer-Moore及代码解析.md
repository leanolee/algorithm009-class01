# Boyer-Moore

#### Brute-Force后缀蛮力匹配算法

```go
/*
Brute-Force：后缀蛮力匹配算法
	1.后缀遍历
	2.匹配成功
	3.i++：BM算法所做的唯一的事情就是改进了这行代码，即模式串不是每次移动一步，而是根据已经匹配的后缀信息，从而移动更多的距离
*/
func bruteForce(s, p string) []int {
	ans, m, n := make([]int, 0), len(s), len(p)
	for i, j := 0, 0; i <= m-n; i++ { // 3
		for j = n - 1; j >= 0 && s[i+j] == p[j]; { // 1
			j--
		}
		if j == -1 { // 2
			ans = append(ans, i)
		}
	}
	return ans
}
```

## Boyer-Moore

#### 坏字符规则

![坏字符规则](E:\极客大学算法训练营第9期\第20课丨字符串算法\坏字符规则.jpg)

1. 模式串存在坏字符b，则模式串的字符b与文本串的坏字符b对齐
2. 模式串不存在坏字符b，则移动模式串到文本串的坏字符b的下一个位置

#### 好后缀规则

![好后缀规则](E:\极客大学算法训练营第9期\第20课丨字符串算法\好后缀规则.jpg)

1. 模式串存在与好后缀完全匹配的子串

   > 最左侧子串与好后缀对齐

   ![好后缀case1](E:\极客大学算法训练营第9期\第20课丨字符串算法\好后缀case1.jpg)

   

2. 模式串存在与好后缀部分匹配的子串

   > 子串（前缀）与好后缀对齐

   ![好后缀case2](E:\极客大学算法训练营第9期\第20课丨字符串算法\好后缀case2.jpg)

   

3. 模式串不存在与好后缀匹配的子串

   > 模式串移动一个身位

   ![好后缀case3](E:\极客大学算法训练营第9期\第20课丨字符串算法\好后缀case3.jpg)

   

#### Code讲解

```go
/*
Boyer-Moore：博伊尔-摩尔算法（后缀字符串匹配）
   p：模式串
   s：文本字符串
   1.坏字符数组的构建
      badChar：[256]int，记录256个（坏）字符对应的挪动位数
      lastIdx：p 的最后一位的索引
      1.1.坏字符规则1：不存在的字符
         全部移动 len(p) 位
      1.2.坏字符规则2：存在的字符
         badChar[i] = lastIdx - i（字符最后一次出现的索引位）
         badChar不记录最后一位字符（lastIdx）
         示例：p = "ABABCB"，lastIdx=5
            A     B     A     B     C     B
            3     2     3     2     1     x
            被覆盖 被覆盖 5-2   5-3   5-4   不记录
   2.好后缀数组的构建
      goodSuffix：记录文本字符串与模式串匹配时，模式串索引 i 位置两者字符不相等，文本字符串需要后移的位数
      prefixIdx：存在与好后缀匹配的前缀时，好后缀的起始索引
         示例：p = "ABABCBAB"
            好后缀 "AB" 存在与之匹配的前缀 "AB"，则 prefixIdx = 6
                      最后一个AB            第一个AB             好后缀AB的起始索引
            第二个 "AB" 不是前缀，前缀必须是从 0 位置开始（但是它符合下面2.3的情况）
      2.1.好后缀规则2：存在与好后缀部分匹配的子串（前缀）
         计算前缀表：
         2.1.1.计算前缀
            示例：p = "ABABCBAB"
               依次比较：前缀 vs p
               "A" vs "B"：lastIdx - i + prefixIdx = 7 - 6 + 8
               "AB" vs "AB"：lastIdx - i + prefixIdx = 7 - 5 + 6
               ...
               "ABABCBA" vs "BABCBAB"：lastIdx - i + prefixIdx = 7 - 0 + 6
            补丁：只有最后一个字符的好后缀，后移位数为 p 的长度 8
         2.1.2.前缀和好后缀匹配，则修正prefixIdx
         2.1.3.修正 goodSuffix[i] 的值
            lastIdx - i：回移到比较的起始位置
            prefixIdx：下一次比较的右移距离（可理解为last+1 - 前缀的长度）
            示例：p = "ABABCBAB"，s = "oooXBABCBABoooooy"，经过计算后的 goodSuffix = [13 12 11 10 9 8 9 8]
               sIdx = 3：oooXBABCBABoooooy
               pIdx = 0：   ABABCBAB
               s[sIdx] != p[pIdx]; sIdx += goodSuffix[pIdx]
                  其中：goodSuffix[pIdx] = 11 = lastIdx - i + prefixIdx
                                               7         0   6
               移动后开始新的一次比较：
               sIdx = 16：oooXBABCBABoooooy
               pIdx = 7：          ABABCBAB
      2.2.好后缀规则3：不存在与好后缀匹配的子串
         同2.1
      2.3.好后缀规则1：存在与好后缀完全匹配的子串
         计算非前缀的完全匹配字符串：i == last，包含在了前缀表的情况
         j：计算完全匹配字符串的长度（即好后缀的长度，suffix len）
         j：回移到比较的起始位置
         lastIdx - i：下一次比较的右移距离
         lastIdx-j：修正后缀表的索引
         2.3.1.计算与好后缀匹配的子串（非前缀）:
            示例：p = "ABABCBAB"，j保存完全匹配字符串的长度（倒序遍历）
               "" vs ABABCBAB：j = 0
               B vs ABABCBAB：j = 1, i = 1, lastIdx - i = 6, lastIdx-j = 6：前缀的情况
               BA vs ABABCBAB：j = 0
               BAB vs ABABCBAB：j = 3, i = 3, lastIdx - i = 4, lastIdx-j = 4：goodSuffix[4] 修正为 7（goodSuffix[lastIdx-j] = j + lastIdx - i）
               BABC vs ABABCBAB：j = 0
               BABCB vs ABABCBAB：j = 1, i = 5, lastIdx - i = 2, lastIdx-j = 6：如 BABCB 最右边的 B 移动到好后缀 B，需要移动 2 个位置，修正 goodSuffix[6] = 1 + 2
               BABCBA vs ABABCBAB：j = 0
         2.3.2.修正goodSuffix：p[i-j] == p[lastIdx-j]：说明是前缀（思考：为什么 if p[i-j] != p[lastIdx-j] { 能排除前缀的情况？）
            示例：p = "ABABCBAB"，s = "oooXBABCBABooooy"，经过计算后的 goodSuffix = [13 12 11 10 7 8 3 1]
               sIdx = 7：oooXBABBBABoooy
               pIdx = 4：   ABABCBAB
               s[sIdx] != p[pIdx]; sIdx += goodSuffix[pIdx]
                  其中：goodSuffix[pIdx] = 7 = j + lastIdx - i
                                              3   7         3
               移动后开始新的一次比较：
               sIdx = 14：oooXBABBBABoooy
               pIdx = 7：        ABABCBAB
   3.根据坏字符数组、好后缀数组，匹配字符串
      3.1.开始匹配
      3.2.匹配成功：重置 s、p 的索引，准备下一次匹配
      3.3.根据坏字符、好后缀规则，挪动 s 的索引
*/
```

#### Golang Code

```go
func boyerMoore(s, p string) []int {
   badChar, goodSuffix := boyerMooreHelper(p) // 1 & 2
   ans, n := make([]int, 0), len(p)-1
   for i := n; i < len(s); { // 3
      j := n
      for ; j >= 0 && s[i] == p[j]; j-- { // 3.1
         i--
      }
      if j < 0 { // 3.2
         i, j = i+1, 0
         ans = append(ans, i)
      }
      i += max(badChar[s[i]], goodSuffix[j]) // 3.3
   }
   return ans
}
func boyerMooreHelper(p string) ([256]int, []int) {
   badChar, goodSuffix, lastIdx := [256]int{}, make([]int, len(p)), len(p)-1
   for i := range badChar { // 1.1
      badChar[i] = len(p)
   }
   for i := 0; i < lastIdx; i++ { // 1.2
      badChar[p[i]] = lastIdx - i
   }
   prefixIdx := len(p)
   goodSuffix[lastIdx] = prefixIdx
   for i := lastIdx - 1; i >= 0; i-- { // 2.1 & 2.2
      j := 0
      for ; j < lastIdx-i; j++ { // 2.1.1
         if p[j] != p[i+1+j] {
            break
         }
      }
      if j == lastIdx-i { // 2.2.2
         prefixIdx = i + 1
      }
      goodSuffix[i] = lastIdx - i + prefixIdx // 2.2.3
   }
   for i := 0; i < lastIdx; i++ { // 2.3
      j := 0
      for ; j < i; j++ { // 2.3.1
         if p[lastIdx-j] != p[i-j] {
            break
         }
      }
      if p[i-j] != p[lastIdx-j] { // 2.3.2
         goodSuffix[lastIdx-j] = j + lastIdx - i
      }
   }
   return badChar, goodSuffix
}
func max(a, b int) int {
   if a > b {
      return a
   }
   return b
}
```

