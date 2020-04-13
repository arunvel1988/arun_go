package main
import "fmt"
func main() {
	var a = make([]int, 4, 7)
	fmt.Printf("Slice1 = %v, \nlength = %d, \ncapacity = %d\n",
			a, len(a), cap(a))
}
