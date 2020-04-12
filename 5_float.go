package main
import "fmt"

func main() {
	a := 5.5
	b := 5.5
	c := a+b
  	d := a-b
	fmt.Printf("Add result is: %f\n", c)
	fmt.Printf("Subtraction result is: %f\n", d)
        fmt.Printf("The data type of d is : %T\n", d)
}
