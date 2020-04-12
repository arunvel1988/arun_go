package main
import "fmt"
func main() {
		nmap := map[int]string {
			1:"hi",
			2:"hello",
		}
		for key, value := range nmap {
			fmt.Println(key, value)
		}
}
