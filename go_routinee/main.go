package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("the first call to hello : ")
	time.Sleep(100 * time.Microsecond)
	fmt.Println("SayHello completed !!")
}

func sayHii() {
	fmt.Println("Say hii executed")
	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println("Learning go routine  :")
	go sayHello()
	go sayHii()

	time.Sleep(100 * time.Microsecond)
}
