package main
import "fmt"
func main() {
	channel1 :=  make(chan int)
	go func(){
		channel1 <- 100
		channel1 <- 200
		channel1 <- 300
		close(channel1)
	}()
	for i := range channel1 {
		fmt.Println(i)
	}
}
