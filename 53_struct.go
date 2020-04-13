package main
import "fmt"

type player struct {
	Name string
	Club string	
	jno int
}
func main() {
	var a player
	fmt.Println(a)
	
	p1 := player{"Messi","Barcelona",10}
	fmt.Println("Player details: ", p1)
	fmt.Println("Player name: ",  p1.Name)
}
