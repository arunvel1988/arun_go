package main
import "fmt"

type player struct {
	Name string
	jno int
}

func main() {
	p1 := &player{"Messi", 10}
	fmt.Println("Player Name: ", (*p1).Name)
 	fmt.Println("Jersey No: ", p1.jno)
}
