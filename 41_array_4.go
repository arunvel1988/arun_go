package main
import "fmt"
func main(){
	a:= [...]int{10,11,12,13,14,15}
	fmt.Println("Length of array :", len(a))


for i:=0; i<len(a); i++{
fmt.Printf("%d\n", a[i])
}
}
