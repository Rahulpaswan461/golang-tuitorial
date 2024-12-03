package main

import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	fmt.Println("Enter your name ")
    reader := bufio.NewReader(os.Stdin)

	name, _ :=reader.ReadString('\n')
	fmt.Println("name is",name)
	var age int
	var address string
     
	fmt.Println("Enter the age and the address")
	fmt.Scanf("%s%d",&age,&address)
	fmt.Println("The age is",age,"the address is",address)
}