package main

import "fmt"

type usuario struct {
	nome string
	idade int8
}

type endereco struct {
	rua string
	cep int32
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Michel"
	u.idade = 24
	fmt.Println(u)

	usuario2 := usuario{"Vidal", 24}	
	fmt.Println(usuario2)

	endereco1 := endereco{"Taquara", 2222222}
	fmt.Println(endereco1)

	endereco2 := endereco{"Curumau", 1111111}
	fmt.Println(endereco2)
}