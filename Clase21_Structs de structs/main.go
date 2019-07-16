package main

import (
	"fmt"
)

type vehiculo struct {
	color   string
	puertas int
}

type camion struct {
	vehiculo
	ruedas bool
}

type ferrari struct {
	vehiculo
	lujoso bool
}

func main() {

	c := camion{
		vehiculo: vehiculo{
			color:   "negro",
			puertas: 4,
		},
		ruedas: true,
	}

	f := ferrari{
		vehiculo: vehiculo{
			color:   "rojo",
			puertas: 4,
		},
		lujoso: true,
	}

	fmt.Println(c)
	fmt.Println(f)
}
