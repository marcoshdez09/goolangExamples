package main

import (
	"fmt"
)

func main() {
	/* f := foo() //es necesario que la funcion que se llame este fuera del main

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f()) */
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)

}

/* func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
} */
func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
