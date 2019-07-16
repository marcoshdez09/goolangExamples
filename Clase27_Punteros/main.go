package main

import (
	"fmt"
)
//Punteros
type persona struct {
	nombre string
}

func main() {

	/* x:=43
	fmt.Println(&x) */

	p1 := persona{
		nombre: "Pedro",
	}

	fmt.Println(p1) //Imprime el valor inicial de la variable
	cambiar(&p1)    //MEtodo que cambia el valor con el puntero
	fmt.Println(p1) //Imprime el valor de la variable cambiada conel puntero
}

func cambiar(p *persona) {//REcibe un valor de tipo 
	p.nombre = "Carlos"
}
