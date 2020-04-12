package main
import "fmt"
func main() {
	a := [5]int{1,2,3,4,5}
	b := a[1:3]
	fmt.Println("Slice :",b)
	fmt.Println("capacity of slice: ", cap(a))
}
