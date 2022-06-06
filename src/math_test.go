package main

import "testing"
import "fmt"

func TestSoma(t *testing.T) {
	total := Soma(11, 15)

	t.Errorf(total)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
