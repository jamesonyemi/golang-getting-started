package main
import "fmt"

func main() {

	num := [] int{42,287,704,25}
	max := num[0]

	for _, value := range num[1:] {
		if max < value {
			max = value
		}
	}
	fmt.Println(max)
}