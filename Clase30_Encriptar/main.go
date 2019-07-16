package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	// "sort"
)

func main() {
	/*
		s := []int{4, 663, 2, 9}
		d := []string{"Juan", "Pedro", "Alejandro"}

		fmt.Println(s)
		sort.Ints(s) //Ordena los valores de tipo int del Slice
		fmt.Println(s)

		fmt.Println(d)
		sort.Strings(d) //Ordena los valores de tipo String del Slice
		fmt.Println(d) */

	//Encriptacion

	s := `contraseniaprueba123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
	
	//Desencriptacion
	clave := `contraseniaprueba123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(clave))
	if err != nil {
		fmt.Println("No te puedes logear")
		return
	}
	fmt.Println("Te has logeado")
}
