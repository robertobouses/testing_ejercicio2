package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio2/user"
)

func TestCalculadoraSuma(t *testing.T) {
	calc := user.Suma{}
	result := calc.Calcular(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Suma incorrecta. Se esperaba %f, se obtuvo %f", expected, result)
	}
}
