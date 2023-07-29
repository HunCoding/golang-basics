package fuzz

import (
	"strings"
)

func MyIndexAny(s, chars string) int {
	/*
		Passa por todos os caracteres da palavra s e verificar
		quantas vezes "chars" aparece dentro dela
	*/
	for i, c := range s {
		if strings.ContainsRune(chars, c) {
			return i
		}
	}

	/*
		Se não existe nenhuma, retorna -1
	*/
	return -1
}
