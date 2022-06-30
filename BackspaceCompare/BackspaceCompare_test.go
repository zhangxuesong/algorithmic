/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/30
 * @Description:
 */
package BackspaceCompare

import "testing"

func TestBackspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "ab#c", t: "ad#c"}, want: true},
		{name: "2", args: args{s: "ab##", t: "c#d#"}, want: true},
		{name: "3", args: args{s: "a#c", t: "b"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("BackspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
