package main
import(
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int)  {
	return a/b, a%b
}

func main() {
	val := add(2,4)
	fmt.Println(val)

	fmt.Println("--------")
  div, mod := divmod(3, 7)
	fmt.Printf("div=%d, mod=%d\n", div, mod)

}