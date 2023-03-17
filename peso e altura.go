package main

import "fmt"

func main() {
	var x1, x2 int
	fmt.Print("Qual o seu peso ?")
	fmt.Scan(&x1)
	fmt.Print("Qual a sua altura")
	fmt.Scan(&x2)
	imc := x1 / (x2 ^ 2)
	fmt.Println("o imc é", imc)
}
