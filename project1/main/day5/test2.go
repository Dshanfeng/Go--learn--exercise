package main

import (
	"fmt"
	"project1/main/day5/structTest"
)

func main() {
	r1 := structTest.Region{3, 3}
	r1.PrintArea('*')

	fmt.Println("----------")

	r2 := structTest.Region{4, 3}
	r2.PrintArea('#')

	fmt.Println("----------")

	r1.PrintMultTable()
}
