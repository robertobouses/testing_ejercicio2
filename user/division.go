package user

import "fmt"

type Division struct{}

func (d Division) Calcular(num1, num2 float64) float64 {
	if num2 != 0 {
		return num1 / num2
	}
	fmt.Println("Error: división por cero.")
	return 0
}
