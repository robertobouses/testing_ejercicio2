package user

type Resta struct{}

func (r Resta) Calcular(num1, num2 float64) float64 {
	return num1 - num2
}
