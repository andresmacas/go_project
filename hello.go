package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scan(&name)
	fmt.Println("Hola", name)
}
