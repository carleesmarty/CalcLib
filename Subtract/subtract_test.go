package Subtract

import (
	"testing"
)

func TestSubtraction(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{

		{
			name: "negative",
			args: args{
				a: -1,
				b: 6,
			},
			want: -7,
		},
		{
			name: "positive",
			args: args{
				a: -1,
				b: -6,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a2 := Subtraction{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
