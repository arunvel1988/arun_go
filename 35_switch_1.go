package main
import "fmt"
func main() {
	
	jersey := 2
	switch {
		case jersey == 1:
		fmt.Println("Ederson")
		case jersey == 2:
		fmt.Println("Ronaldo")
		case jersey == 3:
		fmt.Println("Silva")
		case jersey == 4:
		fmt.Println("Messi")
		default:
		fmt.Println("Invalid")
	}
}
