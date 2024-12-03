package main

import "fmt"

func main(){
	fmt.Println("Learning about the array !!")
	var names[5] int
	names[0] = 1;
	names[1] = 2;
	

	fmt.Println(len(names))
	fmt.Println("the values of the array",names)

	number :=[]int{}
	number = append(number,5)
	fmt.Println(number)

	students := []string{}
	students = append(students,"rahul")
	students = append(students,"paswan")
	fmt.Println("the students are ",students)
}