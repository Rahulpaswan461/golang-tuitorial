package main
import "fmt"
func something(){
	fmt.Println("do something for the given program")
}
func addition(a,b int)(int){
	return a+b;
}
func main(){
	fmt.Println("Hey there we are learning the functions in golang !!")
     something()
	 ans:= addition(4,5)
	 fmt.Println("the addition of two number is", ans)
	 
}