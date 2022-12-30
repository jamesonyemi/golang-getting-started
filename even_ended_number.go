package main
import "fmt"

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			Q := i * j
			S := fmt.Sprintf("%d", Q)
			if S[0] == S[ len(S) - 1 ] {
				count++
			}
		}
	}
	fmt.Println(count)
}