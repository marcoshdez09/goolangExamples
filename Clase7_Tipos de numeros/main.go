package main

import (
	"fmt"
	"runtime"
)

//Tipos de numeros

var z uint32

func main() {

	x := 42
	y := 42.45
	z = 7893

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("Es de tipo %T\n", x)
	fmt.Printf("Es de tipo %T\n", y)
	fmt.Printf("Es de tipo %T\n", z)
	//ver el sistema operativo y la arquitectura
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
