package main
import "fmt"
func main() {
	var a = make([]int, 4, 7)
	for i :=0; i<len(a); i++ {
		fmt.Println(a[i])
	}
}
