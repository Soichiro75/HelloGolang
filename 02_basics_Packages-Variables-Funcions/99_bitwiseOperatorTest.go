package main

import (
	"fmt"
)

var (
	BitwiseOperator uint64 = 10 << 1
)

func main() {
	fmt.Println("BitwiseOperator uint64 = 10<<1")
	fmt.Println("10(Decimal Number) is 1010(Binary Number).")
	fmt.Println("10<<1 means that moves the bits to the left 1, and assigns the rightmost bit a value of 0.")
	fmt.Println("That is means 1010 changes to 10100")
	fmt.Println("So x<<1 means that x make double.")

	fmt.Println("10<<1 means that ", BitwiseOperator) //=>20
}
