package main

import "fmt"

func main() {
	canal := make(chan string, 4)
	canal <- "Ola mundo"
	canal <- "Programando Go"
	canal <- "Programando Go 2"


	mensagem := <- canal	
	mensagem2 := <- canal
	mensagem3 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)

}