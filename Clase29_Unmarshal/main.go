package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
}

func main() {

	s := `[{"Nombre":"Juan","Apellido":"Hernandez"},{"Nombre":"Carlos","Apellido":"Juarez"}]` //Recuperando datos
	bs := []byte(s)                                                                           //Tranformamos a bytes

	fmt.Println(bs)
	var personas []persona               //Declaramos un slice para almacenar los datos
	err := json.Unmarshal(bs, &personas) //Recuperamos datos y alamacenamos en un tipo puntero
	if err != nil {
		fmt.Println(err) //Verdifcamos si hay errores
	}
	for i, v := range personas { //Rescatamos los datos de acuerdo a lo traido

		fmt.Println("La posicion es:", i)
		fmt.Println("Los datos son:", v.Nombre, v.Apellido)

	}

}
