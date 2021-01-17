package SumofTwoIntegers

//不使用运算符+ 和-，计算两整数a、b之和。
//示例 1:
//输入: a = 1, b = 2
//输出: 3
//示例 2:
//输入: a = -2, b = 3
//输出: 1
func GetSum(a int, b int) int {
	for b != 0 {
		s := a ^ b
		b = (a & b) << 1
		a = s
	}
	return a
}