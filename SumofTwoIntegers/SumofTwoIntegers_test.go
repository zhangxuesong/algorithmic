package SumofTwoIntegers

import "testing"

func TestGetSum(t *testing.T) {
	t.Logf("%v \n", GetSum(1,2))
	t.Logf("%v \n", GetSum(-2,3))
	t.Logf("%v \n", GetSum(15,7))
}
