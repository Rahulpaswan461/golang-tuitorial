package main

import "fmt"

func main(){
	day :=30

	switch day{
	case 1,30: 
	    fmt.Println("today is monday")
	case 2:
	    fmt.Println("today is tuesday")
	case 3: 
	     fmt.Println("today is wednesday")
	default:
		fmt.Println("Unknown day")
	}
}