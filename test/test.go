package main

import (
	"fmt"

	piscine ".."
)

func main() {
	a := 1
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
