package main

import (
	"fmt"
)

func main() {
	xi:=[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)//Recorrer elementos de un slice
	fmt.Println("El valor de x", x)
}

func sum(x ...int) int {

	suma := 0
	for i, v := range x {

		suma += v
		fmt.Println("El indice es", i, "El valor que se trae es", v, "y suma es", suma)
	}
	fmt.Println("El total de la suma es:", suma)
	return suma
}
