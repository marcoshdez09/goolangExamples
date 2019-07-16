package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
}

type agente struct {
	persona
	permiso bool
}

//Funcion que retorna
func (s agente) presentar() {

	fmt.Println("Mi nombre es:", s.nombre, s.apellido)

}

func main() {

	s1 := agente{
		persona: persona{
			nombre:   "Jorge",
			apellido: "Hernandez",
		},

		permiso: true,
	}

	fmt.Println(s1)

	//Esta es una funcion anonima
	func(x int) {
		fmt.Println("El valor de x es:", x)
	}(43)

	fmt.Println(pruebas()())//Impresion de la funcion que retorna una funcion

}

//Funcion que retorna una función
func pruebas() func() int {

	return func() int {
		return 1992
	}
}
