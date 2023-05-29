package main

import (
	"fmt"

	"github.com/robertobouses/testing_ejercicio2/user"
)

type Calculadora interface {
	Calcular(num1, num2 float64) float64
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Calculadora Básica")
	fmt.Println("------------------")

	fmt.Print("Ingrese el primer número: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese el segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Ingrese el operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	var calc Calculadora

	switch operator {
	case "+":
		calc = user.Suma{}
	case "-":
		calc = user.Resta{}
	case "*":
		calc = user.Multiplicacion{}
	case "/":
		calc = user.Division{}
	default:
		fmt.Println("Error: operador no válido.")
		return
	}

	result := calc.Calcular(num1, num2)
	fmt.Println("Resultado:", result)
}
