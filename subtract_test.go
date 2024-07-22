package CalcLib

import "testing"

func TestSubtraction(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative",
			args: args{
				a: -1,
				b: 2,
			},
			want: -3,
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
