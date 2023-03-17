package main

import "fmt"

func main() {
	var x1 float64
	fmt.Print("Qual o seu salario ?")
	fmt.Scan(&x1)
	var aumento float64 = 115

	aumento = (x1 * aumento) / 100
	fmt.Println("aumento de ", aumento)
}
