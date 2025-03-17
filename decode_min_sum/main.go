package main

import (
	"fmt"

	"decode_min_sum/logic"
	"decode_min_sum/validate"
)

func main() {
	var encoded string
	fmt.Print("Enter encoded string: ")
	fmt.Scanln(&encoded)

	if validate.CheckConsecutiveLOrR(encoded) {
		// LLLLLLLLLL = 9876543210 ถ้าติดกันมากกว่านี้ จะไม่สามารถ decode ได้
		// RRRRRRRRRR = 0123456789 ถ้าติดกันมากกว่านี้ จะไม่สามารถ decode ได้
		fmt.Println("Cannot decode because there are more than 10 consecutive L or R")
		return
	}

	decoded := logic.DecodeMinSum(encoded)
	fmt.Println("Decoded sequence with minimum sum:", decoded)
}
