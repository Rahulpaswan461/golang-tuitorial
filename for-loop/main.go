package main

import "fmt"

func main(){
	fmt.Println("Learning about the for-loop")
    
	number :=[]int{1,2,3,4,5}
	for i:=0 ; i < len(number); i++{
		fmt.Println("the index is",i,"and the value is",number[i]);
	}

	for index,value :=range number{
		fmt.Println("index is",index,"value is",value)
	}
	data :="rahul paswan"
	for index,value :=range data{
		fmt.Printf("the index and the value is %d and %c\n",index,value)
	}
}