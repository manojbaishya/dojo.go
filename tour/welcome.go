package tour

import "fmt"

func welcome1() {
	fmt.Println("Hello, 世界")
}

func ATourOfGo() {
	fmt.Println("Hello, dojo!")
	fmt.Println("Golang tour...")
	welcome1()
	getterDemo := NewGetterDemo("EntityName", "EntityDescription")
	fmt.Printf("Getter Demo is %v, Name: %s, Description: %s\n", getterDemo, getterDemo.Name, getterDemo.Description)
	fmt.Printf("Getter Demo hash is %d\n", getterDemo.Hash())
}
