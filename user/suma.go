package user

type Suma struct{}

func (s Suma) Calcular(num1, num2 float64) float64 {
	return num1 + num2
}
