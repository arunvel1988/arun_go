	package main
import "fmt"
func main() {
	a :=10

	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	*b = 100
	fmt.Println(*b)
	fmt.Println(a)
}
