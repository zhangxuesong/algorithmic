package ValidParentheses

import "testing"

func TestIsValid(t *testing.T) {
	sMap := make(map[string]bool)
	sMap["()"] = true
	sMap["()[]{}"] = true
	sMap["(]"] = false
	sMap["([)]"] = false
	sMap["{[]}"] = true
	sMap["]"] = false
	for k, v := range sMap {
		b := IsValid(k)
		t.Logf("k= %s, v= %v, b= %v. \n", k, v, b)
	}

	t.Logf("ok")
}
