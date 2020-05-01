package common

import "testing"

func TestTrimSuffixAll(t *testing.T) {
	type args struct {
		val string
		ch  byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Trim Last",
			args: args{
				val: "http://localhost:3000/test/1////",
				ch:  '/',
			},
			want: "http://localhost:3000/test/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSuffixAll(tt.args.val, tt.args.ch); got != tt.want {
				t.Errorf("TrimSuffixAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
