package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("String concepts !!")

	str :="rahul,paswan,dehradun"
	fmt.Println("The value of string",str)

	str2 := strings.Split(str,",")
	fmt.Println("After spliting ", str2)

	str3 :="one two one two three two"

	countOfTwo := strings.Count(str3,"twokk")
	fmt.Println("The count is",countOfTwo)

	str4:="rahul" +" "+ "paswan"
	fmt.Println("the value is",str4)
}
