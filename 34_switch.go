package main
import "fmt"
func main() {
	switch jersey:=3; jersey{
		case 1:
		fmt.Println("Ederson")
		case 2:
		fmt.Println("Ronaldo")
		case 3:
		fmt.Println("Silva")
		case 4:
		fmt.Println("Messi")
		default:
		fmt.Println("Invalid")
	}
}
