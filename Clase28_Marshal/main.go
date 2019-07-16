package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
}

func main() {

	p1 := persona{
		Nombre:   "Jorge",
		Apellido: "Hernandez",
	}
	p2 := persona{
		Nombre:   "Juan",
		Apellido: "Hernandez",
	}
	p3 := persona{
		Nombre:   "James",
		Apellido: "Hernandez",
	}

	personas := []persona{p1, p2, p3}
	fmt.Println(personas)
//Utilizando Marshal con json
	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
