package main

//polimorfismo
import (
	"fmt"
	"math"
)

//Se definen los struct a utilizar de los elementos que van a utilizar
type circulo struct {
	radio float64
}

type cuadrado struct {
	lado float64
}

//Se definen los metodos que utilizaran cada uno y para accesar a las variables de los struct

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c2 cuadrado) area() float64 {
	return c2.lado * c2.lado

}

// se define de manera general que se esta trabajando --enn este caso son formas--
type forma interface {
	area() float64
}

//Funcion para accesae a la informacion de las funciones que saca el area
func informacion(f forma) {
	fmt.Println(f.area())
}
func main() {
	/* //Pasar una funcion a una variable
	f := func() {

		for i := 0; i <= 3; i++ {

			fmt.Println(i)
		}
	}
	f() */
	//Definicion de variables de que llevendatos a los structs
	cir := circulo{
		radio: 12.54,
	}

	cua := cuadrado{
		lado: 5,
	}
	//Funion que muestra la informacion de acuerdo a lo introducido
	informacion(cir)
	informacion(cua)

}
