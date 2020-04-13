package main

import (
	"fmt"
	"sort"
)
func main() {
	a := []int{1,2,3,5,4,3,6,7,90}
	sort.Ints(a)
	fmt.Println("slice: ", a)
}	
