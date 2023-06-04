package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "Variavel 2" // A variavel será string
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Outra forma de declarar variavel
	var (
	variavel3 string = "Variavel3"
	variavel4 bool = true 

	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const nome string = "Michel"
	fmt.Println(nome)
	
}

