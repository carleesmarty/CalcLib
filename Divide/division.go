package Divide

type Division struct{}

func (a2 Division) Calculate(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
