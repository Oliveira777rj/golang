package main

import (
	"fmt"
		
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicação := app.Gerar()
	if erro := aplicação.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}