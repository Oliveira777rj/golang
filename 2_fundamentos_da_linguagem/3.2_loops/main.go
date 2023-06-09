package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 5 {
	// 	i++
	// 	fmt.Println("Incrementando I")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 8; j++ {
	// 	fmt.Println("Incrementação L")
	// 	time.Sleep(time.Second)
	// }

	nomes := [2]string{"Michel", "Vidal"}

	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	carros := [3]string{"Ford", "Chevrolet", "Fiat"}

	for ind, carro := range carros{
		fmt.Println(ind, carro)
	} 
	for indi, letra := range "OLIVEIRA"{
		fmt.Println(indi, string(letra))
	}

	usuario := map[string]string {
		"nome" : "Michel",
		"sobrenome" : "Vidal",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}