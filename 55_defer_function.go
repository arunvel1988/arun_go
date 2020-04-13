package main
import "fmt"

func add (a,b int) int {
	c := a + b
	fmt.Println("Addition: ", c)
	return 0
}

func s1() {
	fmt.Println("hello defer function: ")
}
func main(){
	add(5,15)
	defer add(15,25)
	defer add(25,35)
	s1()
}
