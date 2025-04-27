package basics

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	q, r := divide(77.0, 2.1)
	fmt.Printf("The quotion is: %v, and the reminders is %v\n", q, r)

	val, err := isGreater(6, 6)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

}

func divide(x, y float64) (float64, float64) {
	quotient := x / y
	reminders := math.Mod(x, y)
	return quotient, reminders
}

func isGreater(a, b int) (string, error) {
	if a > b {
		return "A is greater than b", nil
	} else if b > a {
		return "B is greater than a", nil
	} else {
		return "", errors.New("The values is not possible")
	}
}
