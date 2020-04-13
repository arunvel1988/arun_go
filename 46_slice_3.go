package main
import "fmt"
func main(){
	a := []int{1,2,3}
	for index, element := range a{
		fmt.Printf("Index = %d and element = %d\n", index+3, element)
		}
} 
