package user

type Resta struct{}

func (r Resta) Calcular(num1, num2 float64) float64 {
	return num1 - num2 - num2
}

//Cometo un error en la firma de la función
//para que falle el test
