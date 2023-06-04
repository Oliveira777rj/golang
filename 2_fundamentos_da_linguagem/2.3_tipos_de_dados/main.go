package main

import (
	"errors"
	"fmt"
)

/*
Números: Int Float

Int = int8 int16 int32 int64
Float = float32 float64
*/ 

func main() {
	var numero int64 = 5555555555555
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	var numero3 rune = 12345
	fmt.Println(numero3)

	// numero real
	var numeroReal1 float32 = 12.12
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.123
	fmt.Println(numeroReal2)


	// Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	str3 := "Texto3"
	fmt.Println(str3)

	//Boolean
	var booleano1 bool = true
	fmt.Println(booleano1) 

	var booleano2 bool
	fmt.Println(booleano2)


	//Erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)


}