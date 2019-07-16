package main

import (
	"fmt"
)

//Ejemplo de struct
type persona struct {
	nombre   string
	apellido string
	helado   []string
}

func main() {

	p1 := persona{

		nombre:   "Juan",
		apellido: "Hernandez",
		helado: []string{"fresa",
			"Chocolate",
			"vainilla"},
	}
	p2 := persona{

		nombre:   "Jorge",
		apellido: "Hernandez",
		helado: []string{"fresa2",
			"Chocolate2",
			"vainilla2"},
	}

	/* fmt.Println(p1)
	//Imprimir valores especificos de un struct
	for i, v := range p1.helado {

		fmt.Println(i, v)
	}
	for i, v := range p2.helado {

		fmt.Println(i, v)
	} */
	//Imprimir los valores pero accediendo desde un mapa
	m := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}
	//Imprimir valores especificos de un mapa
	//_ en el for no especifica valores de inicializacion
	for _, v := range m {
		fmt.Println(v.apellido)

		// for i, v := range v.helado {

		// 	fmt.Println("", i, v)
		// }
		fmt.Println("............................")
	}
}
