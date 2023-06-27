package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn := "root:Oliveira201&@/filmes_series?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", conn)

	if erro != nil{
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão está aberta com banco de dados")

	linhas, erro := db.Query("SELECT * FROM filmes")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}