package main

import (
	"fmt"

	"dojo/tour"
)

func main() {
	fmt.Println("Hello, dojo!")
	fmt.Println("Golang tour...")
	aTourOfGo()
}

func aTourOfGo() {
	tour.Welcome1()

	getterDemo := tour.NewGetterDemo("EntityName", "EntityDescription")

	fmt.Printf("Getter Demo is %v, Name: %s, Description: %s\n", getterDemo, getterDemo.Name, getterDemo.Description)
	fmt.Printf("Getter Demo hash is %d\n", getterDemo.Hash())
}
