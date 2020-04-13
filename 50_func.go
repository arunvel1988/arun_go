package main
import "fmt"

func f1(a,b int)int{
	c := a+b
	return c
}

func main() {
	fmt.Printf("Addition is %d\n", f1(5,5))
}

	
