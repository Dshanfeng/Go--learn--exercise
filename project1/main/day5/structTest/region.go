package structTest

import "fmt"

type Region struct {
	X, Y int
}

func (region Region) PrintArea(c byte) {
	for i := 0; i < region.Y; i++ {
		for j := 0; j < region.X; j++ {
			print(string(c), " ")
		}
		println()
	}
}

func (Region Region) PrintMultTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
