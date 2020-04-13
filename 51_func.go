package main
import "fmt"
func swap(a,b int)int{
	var temp int
	temp = a
	a = b
	b = temp
 	return temp
}
func main()	{
	var p int = 5
	var q int = 15
		fmt.Printf("p = %d and q = %d", p,q)
	swap(p,q)
		fmt.Printf("\np = %d and q = %d", p,q)

}

