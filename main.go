package main

import "fmt"

// Suma dos números y regresa el resultado
func Sumar(a int, b int) int {
	return a + b
}

func main() {
	resultado := Sumar(3, 5)
	fmt.Println("La suma es:", resultado)
}
