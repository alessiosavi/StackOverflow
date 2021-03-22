package main

import (
	"fmt"
)

func main() {
	var x int
	x = 0
	//var z int
	//result = z

	for {
		fmt.Printf("Geben Sie eine Zahl zum addieren ein: %d\n", x)
		var y int
		_, err := fmt.Scanf("%d\n", &y)
		if err != nil {
			panic(err)
		}
		oldX := x
		x = add(x, y)
		fmt.Printf("SUM: of %d+%d=%d\n", oldX, y, x)
	}
}

func add(x, y int) int {
	return x + y
}
