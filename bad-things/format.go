package main

import "fmt"

func format() {

	hun := "Hun"
	coding := "Coding"

	//Concatenar duas variaveis
	_ = fmt.Sprintf("Concat: %s%#v", hun, coding)

	//Printar duas variaveis e criar uma nova linha
	fmt.Printf("Concat: %s%s\n", hun, coding)
	//ou
	fmt.Println(fmt.Sprintf("Concat: %s%s", hun, coding))

	fmt.Println("Text extenso")

	/*
		Precisava ser assim, sera? Acaba sendo uma complicação que hoje nao faz sentido ter.

		JavaScript
		console.log(`Concat: ${hun + coding}`)

		Rust
		println!("Concat: {}{}", hun, coding)

		Python
		print(f'Concat ${hun + coding}')
	*/
}
