package main
import "fmt"
func main(){
	var x int = 0
	label: for x < 8 {
		if x ==5 {
			x = x + 1
			goto label
		}
		fmt.Printf("value is %d\n",x)
		x++
		}
	}
