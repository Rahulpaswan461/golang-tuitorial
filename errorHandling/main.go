package main

import "fmt"

func division(a,b int)(int,string){
	if(b == 0){

		return 0,"denominator should be greater that 0"
	}

	return a/b,"null"
}


func main(){
	fmt.Println("error handling with functions ")
	ans, _ := division(12,0)
	fmt.Println("Division is",ans)
}