package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Learning about the data conversion")
	num := 12
	fmt.Println("number is", num)
	fmt.Printf("The type of the data is %T\n", num)

	stringVar := strconv.Itoa(num)
	fmt.Println("The value contain by the string is", stringVar)
	fmt.Printf("the type of stringVar %T\n", stringVar)

	name := "rahul_paswan"
	number_name, _ := strconv.Atoi(name)
	fmt.Println("the name is", name)
	fmt.Println("The number name  is ", number_name)
	fmt.Printf("The type is %T\n", number_name)

	piValue := "3.13"
	float_value, _ := strconv.ParseFloat(piValue, 64)
	fmt.Println("the value contains by the float is", float_value)
}
