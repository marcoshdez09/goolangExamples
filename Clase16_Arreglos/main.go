package main

import (
	"fmt"
)

func main() {

	//var x [5]int //Declaracion de un la logitud de un arreglo
	x := []int{1, 2, 3, 4, 5, 6} //Composicion literal de un arreglo--declracion->elementos

	fmt.Println(x)
	fmt.Println(cap(x)) //Define cuantos elementos tiene el arreglo
	//Imprimiendo cada uno de los indices del arreglo
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	//Imprimimos el indice de un arrreglo y sus valores con for
	for i, v := range x {

		println(i, v)
	}
	//imprimir solo un rango de datos de un arreglo
	fmt.Println(x[2:4])
	//Imprimir mismo arreglo con for condicionado por la longitud del mismo
	for s := 0; s < len(x); s++ {
		println(s, x[s])
	}
	//Agragando valores a un slice a otro con append
	x = append(x, 55, 60)
	fmt.Println(x)
	//Segunda forma de agregar slices
	y := []int{555, 443, 233}
	x = append(x, y...)
	fmt.Println(x)
	//Quitar valores especificos de un arreglo
	x = append(x[:2], x[:4]...)
	fmt.Println(x)
	//verificar tamaño de tu arreglo y la capacidad de la misma
	a := make([]int, 10, 100)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
