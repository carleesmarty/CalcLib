package Mod

import "testing"

func TestModulo_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "mod",
			args: args{
				a: 8,
				b: 3,
			},
			want: 2,
		},
		{
			name: "negative",
			args: args{
				a: 6,
				b: -4,
			},
			want: 2,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a2 := Modulo{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
