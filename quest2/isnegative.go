package main

import "github.com/01-edu/z01"

func Isnegative(nbr int) {
	if nbr < 0{
		z01.PrintRune('T')
	} else{
		z01.PrintRune('F')
	}
}

func main() {
	Isnegative(4)
}
