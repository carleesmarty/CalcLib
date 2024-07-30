package Mod

type Modulo struct{}

func (a2 Modulo) Calculate(a, b float64) float64 {

	return float64(int(a) % int(b))
}
