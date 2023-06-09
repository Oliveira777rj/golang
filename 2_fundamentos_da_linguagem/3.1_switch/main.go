package main

import "fmt"

func diaDaSemana(numero int) string{
	switch numero {
	case 1: 
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func main() {
	dia := diaDaSemana(10)
	fmt.Println(dia)

	dia2 := diaDaSemana(1)
	fmt.Println(dia2)
	fmt.Printf(dia2)
}