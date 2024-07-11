package CalcLib

type Calculator interface{ Calculate(a, b int) int }

type Addition struct {
}

func (a2 Addition) Calculate(a, b int) int {
	return a + b

}
