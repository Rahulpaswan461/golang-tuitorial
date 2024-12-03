package main

import "fmt"

type Person struct{
	firstName string
	lastName string
	age int
}

type Contact struct {
	phone string
	call string
}

type Address struct {
	houseNumber int
	street string
	state string
}

type Employee struct {
	personDetails Person
	contactDetails Contact
	address Address
}
func main(){
	fmt.Println("Learning about the struct !!")
	var person1 Person
 
    person1.firstName = "rahulpaswan"
	person1.lastName = "paswan"
	fmt.Println("person1",person1)

	fmt.Println("the information is contain by the function are : ")
	// fmt.Println("the details are ",Employee)

	// var employee1 Employee
    // employee1.personDetails.firstName = "rahul"
	// employee1.personDetails.lastName = "paswan"

	// employee1.contactDetails.phone="12345"
	// employee1.address.houseNumber = 12

	// fmt.Println("the pre-defined datatype is",employee1)
	var employeeDetails Employee
	employeeDetails.personDetails = Person{
		firstName: "rahul",
		lastName: "paswan",
		age: 12,
	}
	fmt.Println("the details are : ", employeeDetails)
}