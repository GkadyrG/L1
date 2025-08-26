package main

import "fmt"

func main() {
	var a, b int

	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка ввода a:", err)
		return
	}

	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка ввода b:", err)
		return
	}

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
