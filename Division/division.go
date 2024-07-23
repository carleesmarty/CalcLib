package Division

type Division struct{}

func (a2 Division) Calculate(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
