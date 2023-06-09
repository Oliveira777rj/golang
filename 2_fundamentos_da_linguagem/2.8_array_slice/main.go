package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Arrays - pouco usado
	var array1[5] int
	array1[0] = 1
	array1[4] = 5
	fmt.Println(array1)

	array2 := [3]string{"posiçao 0", "posiçao 1", "posiçao 2"}
	fmt.Println(array2)

	array2[2] = "posição 2 alterada"
	fmt.Println(array2)

	array3 := [...] float32 {1.1, 1.2, 1.3}
	fmt.Println(array3)

	//Slice
	slice := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(slice)
	fmt.Println(slice[9])
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	// Adicionar valor ao slice = * como o go é fortemente tipado, só é possivel agrupar spice com mesmo tipó
	slice = append(slice, 11)
	fmt.Println(slice)
	

	//Arrays internos
	fmt.Println("........")
	slice3 := make([]float32, 10, 12) 

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 8)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))

}