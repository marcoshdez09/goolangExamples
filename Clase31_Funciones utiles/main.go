package main

import (
	"fmt"
	"sort"
)

type persona struct {
	Nombre string
	Edad   int
}
type PorEdad []persona
type PorNombre []persona

func (a PorEdad) Len() int           { return len(a) } //otiene longitud
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] } //obtiene similares
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }//ordena

func main() {

	p1 := persona{Nombre: "Jorge", Edad: 23}
	p2 := persona{Nombre: "Juan", Edad: 29}
	p3 := persona{Nombre: "James", Edad: 25}

	per := []persona{p1, p2, p3}
	fmt.Println(per) //Imprime d manera desordenada

	sort.Sort(PorEdad(per)) //Implementa eñ metodo para organizar la informacion de acuerdo a la condicion
	fmt.Println(per)

}
