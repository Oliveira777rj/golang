package main

import "fmt"

func somar(n1 int8,n2 int8) int8{
	return n1 + n2
}

func calculoMatematico(m1, m2 int8) (int8, int8) {
	soma := m1 + m2
	subtracao := m1 - m2
	return soma, subtracao
}

func main() {
	soma := somar(2,1)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	f("Texto da função f")

	resultadoSoma, resultadoSubtracao := calculoMatematico(10, 40)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}