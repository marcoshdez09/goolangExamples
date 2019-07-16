package main

import (
	"fmt"
)

func main() {

	m := map[string]int{

		"Hola":  1,
		"Mundo": 2,
	}
	//Imprime los datos del mapa
	fmt.Println(m)
	//Imprime un dato especifico del mapa
	fmt.Println(m["Hola"])
	//Retorna el valor de la variable y el estatus en booleano
	v, ok := m["Go"]

	fmt.Println(v)  //Valor de la variable
	fmt.Println(ok) //Estatus de la variable
	//Condicionando el valor de la variable del mapa si es verdadero
	//por defecto el estatus te lo pone en verdadero
	if v, ok := m["Hola"]; ok {
		fmt.Println("Imprimendo valor de mapa", v)
	}

	m["Go"] = 32 //Agregar nuevos valores al mapa

	//Imprimir todos los valores de un mapa
	for index, v := range m {
		fmt.Println(index, v)
	}
	//Eliminar elementos especificos de un mapa
	delete(m, "Hola")
	fmt.Println(m)
	//Eliminar valores de acuerdo a la condicion del mapa
	if v, ok := m["Go"]; ok {

		fmt.Println("Se elimino el registro con valor:", v) //Mensaje
		delete(m, "Go")                                     //Eliminacion
	}
	fmt.Println(m) //Impresion del mapa nuevamente
}
