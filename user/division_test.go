package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio2/user"
)

func TestCalculadoraDivision(t *testing.T) {
	calc := user.Division{}
	result := calc.Calcular(10, 2)
	expected := 5.0
	if result != expected {
		t.Errorf("División incorrecta. Se esperaba %f, se obtuvo %f", expected, result)
	}
}
