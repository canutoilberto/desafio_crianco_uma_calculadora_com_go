package main

import (
	"math"
	"testing"
)

func ShouldAddCorrect(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("somar(2, 3) = %d; esperado: 5", result)
	}
}

func ShouldSubtractCorrect(t *testing.T) {
	result := subtract(7, 2)
	if result != 5 {
		t.Errorf("subtrair(7, 2) = %d; esperado: 5", result)
	}
}

func ShouldMultiplyCorrect(t *testing.T) {
	result := multiply(3, 4)
	if result != 12 {
		t.Errorf("multiplicar(3, 4) = %d; esperado: 12", result)
	}
}

func ShouldDivideCorrect(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("dividir(10, 2) retornar um erro. %v", err)
	}
	if math.Abs(result-5) > 0.001 {
		t.Errorf("dividir(10, 2) = %f; esperado: 5", result)
	}

	_, err = divide(10, 0)
	if err == nil {
		t.Error("dividir(10, 0) deve retornar um erro.")
	}
}
