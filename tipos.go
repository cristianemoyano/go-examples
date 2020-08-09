package main

import (
	"fmt"
	"strconv"
)

func main()  {
	edad_i := 22

	edad_str := strconv.Itoa(edad_i)
 
	fmt.Println("Mi edad es "+edad_str)

	edad_s := "22"

	edad_int, _ := strconv.Atoi(edad_s)
 
	fmt.Println(edad_int)
}