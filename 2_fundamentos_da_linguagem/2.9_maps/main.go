package main

import "fmt"

func main() {		//valor da chave  // valor do dado	
	usuario := map[string]string {
		"nome": "Michel",
		"sobrenome": "Vidal",
	}
	fmt.Println(usuario["nome"])

	valor := map[int8]string {
		1: "Michel",
		2: "Oliveira",
	}
	fmt.Println(valor)
}