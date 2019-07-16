package main

import (
	"fmt"
)

const (
	_ = iota

	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t%b\t\n", kb, kb)
	fmt.Printf("%d\t%b\t\n", mb, mb)
	fmt.Printf("%d\t%b\t\n", gb, gb)

}
