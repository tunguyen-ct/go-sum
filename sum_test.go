package gosum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			"empty",
			args{a: nil},
			0,
		},
		{

			"1",
			args{a: []int{1}},
			1,
		},
		{

			"2",
			args{a: []int{1, 2}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
