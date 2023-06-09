package main

import "fmt"

type form interface {
	area() float64
}

func escreverArea(f form) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
} 

type retangulo struct {
	altura float32
	largura float32
}

type circulo struct {
	raio float64
}

func main() {
	r := retangulo(10,15)
	escreverArea(r)
}