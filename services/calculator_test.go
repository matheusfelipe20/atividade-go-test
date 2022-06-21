package services_test

import (
	"testing"
	"github.com/matheusfelipe20/atividade-go-test/services"
)

//Ambiente de Teste

func TestSum(t *testing.T){
	if services.Sum(2, 2) != 4 {
		t.Error("2 + 2 deve ser 4")
	}
}

func TestSubtraction(t *testing.T){
	if services.Subtraction(2, 1) != 1 {
		t.Error("2 - 1 deve ser 1")
	}
}

func TestDivision(t *testing.T){
	if services.Division(6, 3) != 2 {
		t.Error("6 / 3 deve ser 2")
	}
}

func TestMultiplication(t *testing.T){
	if services.Multiplication(3, 3) != 9 {
		t.Error("3 * 3 deve ser 9")
	}
}