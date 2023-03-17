package main

import "fmt"

func main() {
	var x1 int
	fmt.Print("qual a sua idade?")
	fmt.Scan(&x1)
	dias := (x1 * 365)
	fmt.Println("a sua idade em dias é", dias)
}
