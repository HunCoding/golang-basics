package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
}

func (p *Pessoa) MarshalText() ([]byte, error) {
	nomeCompleto := fmt.Sprintf("%s %s", p.Nome, p.Sobrenome)
	return []byte(nomeCompleto), nil
}

func (p *Pessoa) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}

	s := string(text)
	parts := strings.Split(s, " ")

	if len(parts) < 2 {
		return errors.New("por favor, especifique primeiro e ultimo nome")
	}

	*p = Pessoa{
		Nome:      parts[0],
		Sobrenome: parts[1],
	}
	return nil
}

func main() {
	var p Pessoa

	valorPadrao := &Pessoa{
		Nome:      "Hun",
		Sobrenome: "Coding",
	}

	flag.TextVar(&p, "pessoa", valorPadrao, "Insira nome e sobrenome desejado")
	flag.Parse()
	fmt.Println(p)
}
