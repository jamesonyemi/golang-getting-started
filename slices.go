package main
import "fmt"

func main() {
	//  same type
	loons := [] string {"daff", "bug", "taz"}
	fmt.Println("loons = %v (type of %T)\n", loons, loons)

	// length
	fmt.Println(len(loons))
	fmt.Println("---------")
	// indexing
	fmt.Println(loons[0])
	fmt.Println("----------")
	// Slices/Slicing
	fmt.Println(loons[1:])

	fmt.Println("-----single value------")
	for i:= range loons {
		fmt.Println(i)
	}

	fmt.Println("-----double value------")
	for i, name:= range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("-----double value, ignore index using underscore, '_'------")
	for _, name:= range loons {
		fmt.Printf("%s\n", name)
	}

	fmt.Println("....Appending to the end of an array object.....")
	loons = append(loons, "elmer")
	fmt.Println(loons)


}
