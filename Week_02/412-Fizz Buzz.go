//写一个程序，输出从 1 到 n 数字的字符串表示。 
//
// 1. 如果 n 是3的倍数，输出“Fizz”； 
//
// 2. 如果 n 是5的倍数，输出“Buzz”； 
//
// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。 
//
// 示例： 
//
// n = 15,
//
//返回:
//[
//    "1",
//    "2",
//    "Fizz",
//    "4",
//    "Buzz",
//    "Fizz",
//    "7",
//    "8",
//    "Fizz",
//    "Buzz",
//    "11",
//    "Fizz",
//    "13",
//    "14",
//    "FizzBuzz"
//]
// 
//
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15
	buzz := fizzBuzz(n)
	fmt.Println(buzz)
}

//leetcode submit region begin(Prohibit modification and deletion)
func fizzBuzz(n int) []string {
	buzz := make([]string, n)
	for i := 0; i < n; i++ {
		num := i + 1
		switch {
		case num%15 == 0:
			buzz[i] = "FizzBuzz"
		case num%3 == 0:
			buzz[i] = "Fizz"
		case num%5 == 0:
			buzz[i] = "Buzz"
		default:
			buzz[i] = strconv.Itoa(i + 1)
		}
	}
	return buzz
}

//leetcode submit region end(Prohibit modification and deletion)
