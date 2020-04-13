package main
import "fmt"
func main(){
	a := []int{1,2,3}
	for _, element := range a{
		fmt.Printf("element = %d\n", element)
		}
} 
