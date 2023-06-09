// go mod init + (nome do modulo) 
package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("michelvidal.com")
	fmt.Println(erro)
}