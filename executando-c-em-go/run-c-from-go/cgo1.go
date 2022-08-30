package main

//#include<stdio.h>
//void inC(int a, int b) {
//    printf("Valor somado: %d", a + b);
//    printf("I am in C code now!\n");
//}
import "C"

import "fmt"

func main() {
	fmt.Println("I am in Go code now!")
	a := C.int(2)
	b := C.int(3)
	C.inC(a, b)
}
