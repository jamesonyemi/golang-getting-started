package main

import "fmt"


func main() {

	for k := 1; k <= 20; k++ {

		if k%3 == 0 {
			fmt.Println("fizz")
		}

		if k%5 == 0 {
			fmt.Println("buzz")
		}

		if ( k%3 == 0 && k%5 == 0) || k%15 == 0 {
			fmt.Println("fizzbuzz")

		}

		if ( k%3 != 0 || k%5 != 0) || k%15 != 0 {
			fmt.Println(k)
		}


	}
}