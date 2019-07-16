package main

import (
	"fmt"
)

func main() {

	defer foo() // con la palabra clave defer pospone el orden en que se ejecuta la funcion
	bar()

}

func foo() {

	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
