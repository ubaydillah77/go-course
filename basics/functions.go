package basics

import "fmt"

func main() {

	fmt.Println(sum(1000, 200))

	// passing a function as an arugment
	hasilnya := gunakanPerhitungan(102, 89, sum)
	fmt.Println("hasilnya", hasilnya)

	// returning and using function
	pangkat2 := buatPerkalian(2)
	fmt.Println("9 * 2 = ", pangkat2(9))

}

func sum(a, b int) int {
	return a + b
}

// function that takes function as argument
func gunakanPerhitungan(x int, y int, perhitungan func(int, int) int) int {
	return perhitungan(x, y)
}

// function that return function
func buatPerkalian(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
