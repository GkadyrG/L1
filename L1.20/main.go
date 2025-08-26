package main

import "fmt"

func reverse(b []rune, i, j int) {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	b := []rune(s)

	reverse(b, 0, len(b)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}

	return string(b)
}

func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Println("Input: ", input)
	fmt.Println("Output:", output)
}
