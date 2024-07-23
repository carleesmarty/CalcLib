package Divide

import "testing"

func TestDivision_Calculate(t *testing.T) {
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
			name: "zero",
			args: args{
				a: -1,
				b: 0,
			},
			want: 0,
		},
		{
			name: "half",
			args: args{
				a: 4,
				b: 2,
			},
			want: 2,
		},
		{
			name: "fractions",
			args: args{
				a: 5,
				b: 2,
			},
			want: 2.5,
		},
		{
			name: "negative",
			args: args{
				a: -4,
				b: 2,
			},
			want: -2,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a2 := Division{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
