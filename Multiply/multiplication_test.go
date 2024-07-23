package Multiply

import "testing"

func TestMultiplication_Calculate(t *testing.T) {
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
				b: 5,
			},
			want: -5,
		},
		{
			name: "double",
			args: args{
				a: 5,
				b: 5,
			},
			want: 25,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a2 := Multiplication{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
