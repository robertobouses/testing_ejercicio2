package user

type Multiplicacion struct{}

func (m Multiplicacion) Calcular(num1, num2 float64) float64 {
	return num1 * num2
}
