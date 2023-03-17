package main

import "fmt"

func main() {
	var x1 int
	fmt.Print("Qual o numero?")
	fmt.Scan(&x1)
	dobro := (x1 * 2)
	fmt.Println("o dobro é", dobro)
	triplo := (x1 * 3)
	fmt.Println("o triplo é", triplo)
	quadruplo := (x1 * 4)
	fmt.Println("o quadruplo é", quadruplo)
}
