package main

type Valores struct {
	x int //nao precisa ter virgula
	y int
}

func inverterValores(x, y int) (int, int) { //aqui sao obrigatorios os parenteses
	return y, x //mas aqui nao
}

func syntaxExample() {
	_ = Valores{
		x: 50,
		y: 100, // por que precisa de virgula no ultimo parametro?
	}
}
