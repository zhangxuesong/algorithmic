/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/23
 * @Description:
 */
package MoveZeroes

import "testing"

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "a", args: args{nums: []int{0, 1, 0, 3, 12}}},
		{name: "b", args: args{nums: []int{0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.args.nums)
		})
	}
}
