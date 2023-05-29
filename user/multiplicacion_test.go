package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio2/user"
)

func TestCalculadoraMultiplicacion(t *testing.T) {
	calc := user.Multiplicacion{}
	result := calc.Calcular(4, 10)
	expected := 40.0
	if result != expected {
		t.Errorf("Multiplicación incorrecta. Se esperaba %f, se obtuvo %f", expected, result)
	}
}
