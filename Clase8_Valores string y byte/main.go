package main

import (
	"fmt"
)

func main() {

	s1 := "hola mundo"
	s2 := `linea de string`

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//Creacion de ciclo for para ver los valores UTF8
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U\n", s1[i])
	}

}
