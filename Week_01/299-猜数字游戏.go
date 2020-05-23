//你正在和你的朋友玩 猜数字（Bulls and Cows）游戏：你写下一个数字让你的朋友猜。每次他猜测后，你给他一个提示，告诉他有多少位数字和确切位置都猜对
//了（称为“Bulls”, 公牛），有多少位数字猜对了但是位置不对（称为“Cows”, 奶牛）。你的朋友将会根据提示继续猜，直到猜出秘密数字。 
//
// 请写出一个根据秘密数字和朋友的猜测数返回提示的函数，用 A 表示公牛，用 B 表示奶牛。 
//
// 请注意秘密数字和朋友的猜测数都可能含有重复数字。 
//
// 示例 1: 
//
// 输入: secret = "1807", guess = "7810"
//
//输出: "1A3B"
//
//解释: 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。 
//
// 示例 2: 
//
// 输入: secret = "1123", guess = "0111"
//
//输出: "1A1B"
//
//解释: 朋友猜测数中的第一个 1 是公牛，第二个或第三个 1 可被视为奶牛。 
//
// 说明: 你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。 
// Related Topics 哈希表
package main

import (
	"fmt"
)

func main() {
	//secret := "1807"
	//guess := "7810"
	//secret := "1123"
	//guess := "0111"
	//secret := "1122"
	//guess := "1222"
	//secret := "11"
	//guess := "10"
	secret := "011"
	guess := "110"
	hint := getHint(secret, guess)
	fmt.Println(hint)
}

/*
暴力法：O(n^3),O(n)
	三层嵌套循环：逻辑较复杂，且判重不方便
	1.使用memo记录已经被映射过的 secret 的索引
	2.一次循环，对比 secret 和 guess 相同的索引
		值相同，公牛+1，并将索引存入 memo
	3.二次循环：
		如果是 公牛+1 过的位置，跳过
		开始循环遍历 secret：两边索引相同 或 两边值不同，直接跳过
			值相同：
			1.遍历memo，索引存在过，终止本次遍历 secret 的循环
			2.没存在过，奶牛+1
hash映射：O(n),O(n)
	1.一次循环，对比 secret 和 guess 相同的索引
		值相同，公牛+1；值不同，hash中存 k:值，v:出现的次数
	2.二次循环，不是公牛的情况下，从hash中取 k,v，k存在且 v>0，母牛+1
hash映射：优化，O(n),O(n)
	仅一次循环：
		使用数组缓存出现过的次数，存对应 0-10/256个ascii 字符
		相同索引值相同，公牛+1
		不同索引
			secret 的字符出现，缓存中 +1。如果+1前，值小于0，奶牛+1
			guess 的字符出现，缓存中 -1。如果-1前，值大于0，奶牛+1
*/
//leetcode submit region begin(Prohibit modification and deletion)
func getHint(secret string, guess string) string {
	// leetcode一次循环
	bulls, cows := 0, 0
	nums := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			if nums[secret[i]-'0'] < 0 {
				cows++
			}
			nums[secret[i]-'0']++
			if nums[guess[i]-'0'] > 0 {
				cows++
			}
			nums[guess[i]-'0']--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)

	// 个人方法
	//bulls, cows := 0, 0
	//memo := make(map[byte]int)
	//for i := 0; i < len(secret); i++ {
	//	if secret[i] == guess[i] {
	//		bulls++
	//	} else {
	//		memo[secret[i]]++
	//	}
	//}
	//for i := 0; i < len(secret); i++ {
	//	if secret[i] != guess[i] {
	//		if b, ok := memo[guess[i]]; ok && b > 0 {
	//			memo[guess[i]]--
	//			cows++
	//		}
	//	}
	//	//if secret[i] != guess[i] {
	//	//	if b, ok := memo[guess[i]]; ok {
	//	//		if b > 0 {
	//	//			memo[guess[i]]--
	//	//			cows++
	//	//		}
	//	//	}
	//	//}
	//}
	//return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"

	// 暴力法
	//bulls, cows := 0, 0
	//memo := make([]int, 0)
	//for i := 0; i < len(guess); i++ {
	//	// 要先检查相同位置
	//	if guess[i] == secret[i] {
	//		bulls++
	//		memo = append(memo, i)
	//	}
	//}
	//for i := 0; i < len(guess); i++ {
	//	if guess[i] == secret[i] {
	//		continue
	//	}
	//	for j := 0; j < len(secret); j++ {
	//		if i == j || guess[i] != secret[j] {
	//			continue
	//		}
	//		for k := 0; k < len(memo); k++ {
	//			if memo[k] == j {
	//				goto secret
	//			}
	//		}
	//		cows++
	//		memo = append(memo, j)
	//		break
	//	secret:
	//	}
	//}
	//return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

//leetcode submit region end(Prohibit modification and deletion)
