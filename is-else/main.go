package main
import "fmt"
func main(){
	x := 20

	if x == 2 {
		fmt.Println("the number is under the range")
	}else if x == 10 {
		fmt.Println("the number is equal to 10")
	}else{
		fmt.Println("Number is smaller i guess")
	}
}