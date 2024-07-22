package CalcLib

import "testing"

func TestAddition(t *testing.T) {
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
			name: "zeros",
			args: args{
				a: 0,
				b: 0,
			},
			want: 0,
		},

		{
			name: "negative",
			args: args{
				a: -1,
				b: 0,
			},
			want: -1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a2 := Addition{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
