package TwoSum

import "testing"

func TestTwoSum(t *testing.T) {
	t.Logf("结果是：%v \n", TwoSum([]int{2, 7, 11, 15}, 9))
	t.Logf("结果是：%v \n", TwoSum([]int{3, 2, 4}, 6))
	t.Logf("结果是：%v \n", TwoSum([]int{3, 3}, 6))
}
