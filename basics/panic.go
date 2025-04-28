package basics

import "fmt"

func main() {
	// panic(interface{})
	process(9999)

	process(-1)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input harus lebih dari 0")
	}

	fmt.Println("Processing input", input)
}
