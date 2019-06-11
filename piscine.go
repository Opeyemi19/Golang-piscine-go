package piscine

import (
	"strconv"
)

// Atoi ...
func Atoi(s string) int {

	// if s == ParseInt() {
	// 	// return 0
	// }

	if S, err := strconv.Atoi(s); err == nil {
		// fmt.Printf("%v", S)
		// fmt.Println(S)
		return S
	}

	return 0

}
