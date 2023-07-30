package main

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	s := 0
	PointOne(&s)
	fmt.Println(s)
}
