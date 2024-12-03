package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning about time package ::")
	time := time.Now()
	fmt.Println("the current time is ", time)
	formattedTime := time.Format("2006/01/02, Monday")
	fmt.Println("formatted time ", formattedTime)
}
