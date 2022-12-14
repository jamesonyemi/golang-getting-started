package main

import("fmt")

func main(){
	y := 100

	switch y {
	case 1:
		fmt.Println("not found")
	case 100:
		fmt.Printf("found:%v", y)
	default:
		fmt.Println("belong not here")
	}
}