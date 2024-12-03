package main

import "fmt"

func modifyTheValueOfThevaraible(num *int){
	*num = *num  + 2

}

func main(){
	fmt.Println("Learning about the pointers !!")
	// num := 10
	// fmt.Println("the value contains by the variable is", num)
	// var ptr *int
	// ptr = &num
	// fmt.Println("the address contains  by the pointer", ptr)
	// fmt.Println("the exact value contains by the poitner",*ptr)

	number := 10
	modifyTheValueOfThevaraible(&number)
	fmt.Println("the modify value of the num is", number)
}