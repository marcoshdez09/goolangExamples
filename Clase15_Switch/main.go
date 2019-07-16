package main

//Se hace la comparacion en cada una de las opciones y en cada opcion puede validar dos opciones
func main() {

	switch { //se puede poner la primera declaración  para compara con los de mas casos, puede ser otros tipos de valores ya sea string o booleanos

	case 2 == 2:
		println("Primera opcion")

	case 4 == 8:
		println("Es la segunda opcion")

	case 7 < 8:
		println("Es la tercera opcion")

	default:
		println("Es un valor por default")

	}
}
