package main

import "fmt"

func main(){
	fmt.Println("Learning about the map")

	studentsGrades := make(map[string]int)
	studentsGrades["rahul" ] = 1000
	studentsGrades["paswan"] = 200
	studentsGrades["sanjay"] = 120
	studentsGrades["ram"] = 1222


	fmt.Println("the makrs of the rahul is", studentsGrades["rahul"])

	delete(studentsGrades,"sanjay")

	fmt.Println("the value of the sanjay is",studentsGrades["sanjay"])

	grades, exists := studentsGrades["rahul"]

	fmt.Println("grades is",grades,"and the person is exists or not",exists)

	for index,value := range studentsGrades{
		fmt.Println("index is",index,"value is",value)
	}
}