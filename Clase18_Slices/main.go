package main

import (
	"fmt"
)

func main() {

	//slice de slices
	x := []string{"Hola", "Mundo", "Go"}
	x_ := []string{"Como", "Clase18"}

	xs := [][]string{x, x_}

	fmt.Println(xs)

	for i, r := range xs {

		fmt.Println("EL indice es", i)
		for j, val := range r {
			fmt.Println(j, val)
		}
	}
	//Crear mapa de slices
	a := map[string][]string{

		"Jorge": []string{"Futbol", "Pizza", "Cine"},
		"Fabi":  []string{"Videojuegos", "Pizza", "Cine"},
		"Rich":  []string{"Nada", "Pizza", "Cine"},
	}
	a["Abi"] = []string{"Games", "Pollo", "Parque"} //Forma de agregar un elemento en el mapa de slices de string
	delete(a, "Abi")                                //Forma de eliminar un ragistro de un mapa
	for z, ra := range a {

		fmt.Println(z)
		for g, ra_ := range ra {
			fmt.Println(g, ra_)
		}
	}

	//mapa de slices de int

	b := map[int][]int{

		1: []int{1, 2, 3, 4, 5},
		2: []int{6, 7, 8, 9, 10},
		3: []int{11, 12, 13, 14, 15},
	}

	b[4] = []int{16, 17, 18, 19, 20} //Forma de agregar registros en un mapa de slices de int
	delete(b, 2)                     //Forma de eliminar un elemento del mapa
	for d, rint := range b {
		fmt.Println(d)

		for f, r_int := range rint {

			fmt.Println(f, r_int)
		}
	}

}
