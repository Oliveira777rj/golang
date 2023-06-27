package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type estudante struct {

	Nome string `json:"Nome"`
	Sobrenome string `json:"Sobrenome"`
	Habilidades string `json:"Habilidades"`
	Idade uint `json:"Idade"`
}

func main() {
	c := estudante{"Silvia", "Olivia", "Node", 20}
	fmt.Println(c)

	estudanteEmJSON, erro := json.Marshal(c) 
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(estudanteEmJSON))

	c2 := estudante{"Maria", "Alvarez", "Java", 35}
	estudante2EmJSON, erro := json.Marshal(c2)
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(estudante2EmJSON))
}