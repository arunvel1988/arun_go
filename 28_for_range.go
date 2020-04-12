package main
import "fmt"
func main() {
	a := []string{"hi", "hello", "worlfd"}
	for i, j := range a {
		fmt.Println(i, j)
	}
}
