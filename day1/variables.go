package day1

import (
	"fmt"
)

func VaribaleDeclaration() {
	var a int
	var b int = 10
	var c string = "I am String"
	e := "I am e"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)

	// variable shadowing
	// avoid this type of code
	if true {
		var b int = 16
		e := "changed value from e to f"
		fmt.Printf("shadowed value of %s \n", b)
		fmt.Println(e)

	}
}
