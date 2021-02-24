package FibonacciNumber

import (
	"fmt"
	"testing"
)

func TestFibonacciNumber(t *testing.T) {
	num := 11
	res := FibonacciNumber(num)
	fmt.Println(res)
}
