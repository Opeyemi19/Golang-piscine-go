package piscine

import "fmt"

//IsNegative ...
func IsNegative(nb int) {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}
