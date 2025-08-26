package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	var aStr, bStr, op string

	fmt.Print("Введите число a: ")
	fmt.Scan(&aStr)

	fmt.Print("Введите число b: ")
	fmt.Scan(&bStr)

	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)

	a, ok := new(big.Int).SetString(aStr, 10)
	if !ok {
		log.Fatal("Ошибка: не удалось преобразовать число a")
	}

	b, ok := new(big.Int).SetString(bStr, 10)
	if !ok {
		log.Fatal("Ошибка: не удалось преобразовать число b")
	}

	result := new(big.Int)

	switch op {
	case "+":
		result.Add(a, b)
	case "-":
		result.Sub(a, b)
	case "*":
		result.Mul(a, b)
	case "/":
		if b.Cmp(big.NewInt(0)) == 0 {
			log.Fatal("Ошибка: деление на ноль")
		}
		result.Div(a, b)
	default:
		log.Fatalf("Ошибка: неизвестная операция %s", op)
	}

	fmt.Println("Результат:", result)
}
