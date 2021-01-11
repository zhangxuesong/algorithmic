package ValidParentheses

//有效的括号
//https://leetcode-cn.com/problems/valid-parentheses/
func IsValid(s string) bool {
	//计算字符串长度
	n := len(s)
	//括号一定是成对的，字符串长度一定是偶数
	if n%2 == 1 {
		return false
	}
	//声明右括号作为key的字典，方便后面配对
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	//声明byte数组作为存储待配对字符的栈
	stack := []byte{}
	//遍历字符串
	for i := 0; i < n; i++ {
		//判断左右括号
		if pairs[s[i]] > 0 {
			//key在pairs中存在表示当前字符是右括号
			//stack为0，表示没有可以和当前字符配对的左括号
			//判断stack最后一位是否和当前字符成对
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			} else {
				//配对成功
				//删除数组末尾元素表示出栈
				stack = stack[:len(stack)-1]
			}
		} else {
			//key在pairs中不存在表示当前字符是左括号
			//数组中添加元素表示压栈
			stack = append(stack, s[i])
		}
	}
	//stack长度为空说明全部配对成功，否则不成功
	return len(stack) == 0
}
