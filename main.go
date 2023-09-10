package main

import "fmt"

func main() {
	s := []int{0, 1, 1, 1, 2, 2, 3, 4, 4}
	fmt.Println(s)

	UniqInPlaceSorted(&s)
	fmt.Println(s)
}
