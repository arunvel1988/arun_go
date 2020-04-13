package main
import "fmt"

type player struct {
	name string
	jno int
}

func (p1 player) show() {
	fmt.Println("Player name: ", p1.name)
	fmt.Println("Jersey No: ", p1.jno)
}

func main() {
	a :=  player{
	name: "Messi",
	jno: 10,
	}

	a.show()
}


