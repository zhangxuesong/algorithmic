package FizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	n := 15
	res := FizzBuzz(n)
	t.Logf("the res is: %v \n", res)
}
