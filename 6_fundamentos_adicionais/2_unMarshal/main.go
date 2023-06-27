package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type estudante struct {

	Nome string `json:"Nome"`
	Sobrenome string `json:"Sobrenome"`
	Habilidades string `json:"Habilidades"`

}

func main() {
	estudanteEmJSON := `{"Nome":"Maria","Sobrenome":"Alvarez","Habilidades":"Java"}`

	var e estudante 

	if erro := json.Unmarshal([]byte(estudanteEmJSON),&e); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(e)

	estudante2EmJSON := `{"Nome":"toby", "Sobrenome":"Vieira", "Habilidades":"Python"}`

	e2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(estudante2EmJSON), &e2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(e2)
}