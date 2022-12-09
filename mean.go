package main

import(
	"fmt"
)

func main() {

	x := 1.0
	y := 3.00

	fmt.Printf("x=%v, of type %T\n", x, x)
	fmt.Printf("y=%v, of type %T\n", y, y)

	mean := (x+y)/2
	fmt.Printf("the mean: %v, type of %T\n", mean, mean)

}