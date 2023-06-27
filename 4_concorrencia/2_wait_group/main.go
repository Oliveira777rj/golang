package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Ola mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando go")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("go routine 3")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("go routine 4")
		waitGroup.Done() // -1
	}()
	waitGroup.Wait() // Espera o contador chegar a 0
}


func escrever(texto string) {
	for i := 0; i < 2; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}