package main
import "fmt"

func main() {
	s1 := "hello"
	s2 := "hello"
	s3 := "Hello"

	s4 := s1 == s2
	s5 := s2 == s3
	fmt.Println(s4)
	fmt.Println(s5)
        fmt.Printf("string length is %d\n", len(s1))
	fmt.Printf("The type of s4 is %T\n", s4)
}
