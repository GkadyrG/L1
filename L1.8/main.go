package main

import (
	"fmt"
)

func SetBit(n int64, i uint, value bool) int64 {
	if value {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func main() {
	var n int64 = 5
	fmt.Printf("Исходное число: %d (бинарно: %04b)\n", n, n)

	n1 := SetBit(n, 1, false) // установить 1-й бит в 0
	fmt.Printf("После установки 1-го бита в 0: %d (бинарно: %04b)\n", n1, n1)

	n2 := SetBit(n, 2, true) // установить 2-й бит в 1
	fmt.Printf("После установки 2-го бита в 1: %d (бинарно: %04b)\n", n2, n2)

	n3 := SetBit(n, 3, true) // установить 3-й бит в 1
	fmt.Printf("После установки 3-го бита в 1: %d (бинарно: %04b)\n", n3, n3)
}
