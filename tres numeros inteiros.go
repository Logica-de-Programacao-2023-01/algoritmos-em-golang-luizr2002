package main

import "fmt"

func main() {
	var x1, x2, x3 int
	fmt.Print("Qual o valor do primeiro numero")
	fmt.Scan(&x1)
	fmt.Print("Qual o valor do segundo numero")
	fmt.Scan(&x2)
	fmt.Print("Qual o valor do terceiro numero")
	fmt.Scan(&x3)
	soma := (x1 + x2 + x3)
	fmt.Println("a some é", soma)
}
