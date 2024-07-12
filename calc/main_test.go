package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/carleesmarty/CalcLib"
)

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
			a2 := CalcLib.Addition{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerHappyPath(t *testing.T) {
	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1", "4"})

	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(calculator.inputs, []int{1, 4}) {
		t.Errorf("Incorrect inputs: %v", calculator.inputs)
	}
	if buffer.String() != "42" {
		t.Errorf("Incorrect outputs: %s", buffer.String())
	}
}

func TestHandlerArgumentsFail(t *testing.T) {

	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1"})
	if err == nil {
		t.Errorf("Incorrect inputs: %v", calculator.inputs)
	}
}

func TestHandlerArgumentsPass(t *testing.T) {
	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1", "3"})
	if err != nil {
		t.Error(err)
	}
}

type fakeCalculator struct {
	inputs []int
	output int
}

func (f *fakeCalculator) Calculate(a, b int) int {
	f.inputs = append(f.inputs, a, b)
	return f.output
}
