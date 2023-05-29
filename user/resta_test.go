package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio2/user"
)

func TestCalculadoraResta(t *testing.T) {
	calc := user.Resta{}
	result := calc.Calcular(5, 3)
	expected := 2.0
	if result != expected {
		t.Errorf("Resta incorrecta. Se esperaba %f, se obtuvo %f", expected, result)
	}
}
