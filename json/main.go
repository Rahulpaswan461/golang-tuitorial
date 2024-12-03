package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	IsAdult bool   `json:"IsAdult`
}

func main() {
	fmt.Println("learning about json")
	person := Person{Name: "Rahul", Age: 21, IsAdult: true}
	fmt.Println("person details is", person)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	fmt.Println("the json data is", string(jsonData))

	// unmarshal the data
	var personsData Person
	err = json.Unmarshal(jsonData, &personsData)
	if err != nil {
		fmt.Println("the error is", err)
		return
	}
	fmt.Println("person data after unmarshal", personsData)

}
