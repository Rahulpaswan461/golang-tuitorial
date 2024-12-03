package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TODO struct {
	Userid    int    `json:"user_id"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"Completed"`
}

func performGetRequest() {
	todos := TODO{
		Userid:    2,
		Title:     "Rahul Paswan",
		Completed: true,
	}
	// convert the todo struct to json data
	jsonData, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("Error while marshaling", err)
		return
	}
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error while sendin the data", err)
		return
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response data is", res.StatusCode, "and data is", string(data))
}
func getSomeDataAndCreate() {
	todos := TODO{
		Userid:    1,
		Title:     "Dehradun",
		Completed: true,
	}

	jsonData, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("Error while parsing data")
		return
	}
	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while creating data", err)
		return
	}
	fmt.Println("response code is", res.StatusCode)

}

func performDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error creating request ", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status", res.StatusCode)
}

func main() {
	//get request operations are there
	// fmt.Println("Designing the APIS")
	// res, err := http.Get("https://jsonplaceholder.typicode.com/todo")
	// if err != nil {
	// 	fmt.Println("Error while fetching the data", err)
	// 	return
	// }
	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println("Bad request from the frontend")
	// 	return
	// }
	// var todos TODO
	// err = json.NewDecoder(res.Body).Decode(&todos)

	// if err != nil {
	// 	fmt.Println("Error while parsing the information !!")
	// 	return
	// }
	// fmt.Println("the values of the todos are : ", todos)
	// fmt.Println("the values are ", todos.Description)
	// fmt.Println("Title is ", todos.Title)
	performGetRequest()
	getSomeDataAndCreate()
	performDeleteRequest()

}
