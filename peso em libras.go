package main

import "fmt"

func main() {
	var x1 int
	fmt.Print('Qual o seu peso em Kg ?')
	fmt.Scan(&x1)
	libras := (x1 * 220) / 100
	fmt.Println("seu peso em libras é", libras)
}
