package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	results := make([]int, len(numbers))
	
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for i, num := range numbers {
		go func(index int, n int) {
			defer wg.Done() 
			
			results[index] = n * n
		}(i, num) // Исправлено в версиях Go 1.22+
	}
	
	wg.Wait() 
	
	fmt.Println("Результаты в порядке массива:")
	for i, num := range numbers {
		fmt.Printf("Квадрат числа %d: %d\n", num, results[i])
	}
}
