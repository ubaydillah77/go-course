package main

import "fmt"

func main() {
	// ... Ellipsis
	// func functionName(param1 type1, param2 ...type2) returnType{
	// }
	// variadic function must be the last parameter

	statements, total := sum("The result of sum is =", 5, 122, 55)
	fmt.Println(statements, total)

}

func sum(returnString string, nums ...int) (string, int) {
	total := 0

	for _, v := range nums {
		total += v

	}
	return returnString, total
}
