package main

import (
	"fmt"
)

//Creando Structs que permiten a la vez crear tipos de datos
type persona struct {
	nombre   string
	apellido string
	edad     int
	calle    string
	casanum  string
}

//Para Struct embebido
type general struct {
	persona

	verifica bool
}

func main() {
	//Agregando datos
	persona1 := persona{
		nombre:   "Juan",
		apellido: "Hernandez",
		edad:     14,
		calle:    "Benito Juarez",
		casanum:  "12",
	}
	persona2 := persona{
		nombre:   "Jorge",
		apellido: "Hernandez",
		edad:     12,
		calle:    "Florida",
		casanum:  "9",
	}
	//Struct embebido
	persona3 := general{

		persona: persona{
			nombre:   "Ricardo",
			apellido: "Juarez",
			edad:     4,
			calle:    "Ninguna",
			casanum:  "0",
		},
		verifica: true,
	}
	//Struct anonimo
	anonimo := struct {
		nombre   string
		apellido string
	}{
		nombre:   "Emanuel",
		apellido: "Hernandez",
	}

	fmt.Println(anonimo) //Imprime el struct anonimo
	//Imprime el struct embebido
	fmt.Println(persona3.nombre, persona3.apellido, persona3.edad, persona3.calle, persona3.casanum, persona3.verifica)
	//Imprime los structs normales
	fmt.Println(persona1.nombre, persona1.apellido, persona1.edad, persona1.calle, persona1.casanum)
	fmt.Println(persona2.nombre, persona2.apellido, persona2.edad, persona2.calle, persona2.casanum)

	// fmt.Println(persona1)

}
