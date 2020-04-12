package main
import "fmt"
func main() {
	var a[2][2] int
	a[0][0] = 10
	a[0][1] = 11
	a[1][0] = 12
	a[1][1] = 13
	
	for i:=0; i<2; i++{
	for j:=0; j<2; j++{
	fmt.Println(a[i][j])
	}
	
	}
}
