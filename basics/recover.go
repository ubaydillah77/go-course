package basics

import "fmt"

func main() {

	process()

	fmt.Println("Melanjutkan process")
}

func process() {
	defer func() {
		// if r:= recover(); r != nil
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("mulai process")
	panic("Sesuatu berjalan tidak benar!")
	fmt.Println("Selesaikan process")
}
