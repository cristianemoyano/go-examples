package main

import (
	"fmt"
	"unsafe"
)

func main() {
	num1 := 5
	num2 := 7
	num3 := &num1
	num4 := &num2
	fmt.Println(*num3 + *num4)
	fmt.Println(&num1)
	fmt.Println(&num2)
	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
}