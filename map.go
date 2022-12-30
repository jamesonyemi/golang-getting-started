package main
import "fmt"

func main() {
	stock := map[string] float64 {
		"AMZN": 299.9,
		"GOOG": 120.5,
		"MSFT": 978,
	}

	// number of items
	fmt.Println(len(stock))

	for a, value:= range stock {
		fmt.Printf("%s: %v\n", a, value)
	}

	fmt.Println("--get a value---")
	fmt.Println(stock["AMZN"])

	fmt.Println("--get zero value, if the value is not found")
	fmt.Println(stock["Jumia"])

	fmt.Println(`
	--use two value to see if
	a given value is not found
	`,
)
 value, isOk := stock["TSLA"]
 if  !isOk {
 fmt.Println("found TSLA")
 } else {
	fmt.Println(value)
}

fmt.Println("--set values")
stock["TSLA"] = 233.09
fmt.Println(stock)



}