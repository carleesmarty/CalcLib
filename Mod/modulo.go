package Mod

type Modulo struct{}

func (a2 Modulo) Calculate(a, b int) float64 {

	return float64(a % b)
}
