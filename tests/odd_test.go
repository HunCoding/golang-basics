package tests

import "testing"

// Test
func TestIsOdd_return_odd(t *testing.T) {
	// Definindo varias
	var value int64 = 3

	//Chamar as funcoes necessarios
	number_type := IsOdd(value)

	//Verificar se o resultado foi conforme esperado
	if number_type != ODD_KEYWORD {
		t.Errorf("%s - expected %s, received %s",
			t.Name(),
			ODD_KEYWORD,
			number_type,
		)
		return
	}
}

func TestIsOdd_return_even(t *testing.T) {
	// Definindo varias
	var value int64 = 4

	//Chamar as funcoes necessarios
	number_type := IsOdd(value)

	//Verificar se o resultado foi conforme esperado
	if number_type != EVEN_KEYWORD {
		t.Errorf("%s - expected %s, received %s",
			t.Name(),
			EVEN_KEYWORD,
			number_type,
		)
		return
	}
}
