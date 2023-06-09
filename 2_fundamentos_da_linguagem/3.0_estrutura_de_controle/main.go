package main

import "fmt"

func main() {
	numero := 0

	if numero > 40{
		fmt.Println("Maior que 15")
	}else {
		fmt.Println("Menor ou igual a 15")
	}
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}else if numero < 0{
		fmt.Println("O numero é menor que zero")
	}else{
		fmt.Println("O numero e igual a zero")
	}
}