package FizzBuzz

import "strconv"

//写一个程序，输出从 1 到 n 数字的字符串表示。
//1. 如果n是3的倍数，输出“Fizz”；
//2. 如果n是5的倍数，输出“Buzz”；
//3.如果n同时是3和5的倍数，输出 “FizzBuzz”。
func FizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 && i%5 == 0 {
			s = "FizzBuzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else {
			s = strconv.Itoa(i)
		}
		res = append(res, s)
	}
	return res
}
