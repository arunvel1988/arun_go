package main
import "fmt"
func main() {
	a := 5
	b := 10
	
	b = a
	fmt.Println(a)
	fmt.Println(b)
        a--  // a= a+b
	fmt.Println(a)
}
