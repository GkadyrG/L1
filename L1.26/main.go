package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {
	seen := make(map[rune]struct{})
	lowerStr := strings.ToLower(s)

	for _, char := range lowerStr {
		if _, exists := seen[char]; exists {
			return false
		}
		seen[char] = struct{}{}
	}
	return true
}

func main() {
	tests := []string{
		"abcd",     // true
		"bCdefAaf", // false
		"aabcd",    // false
		"",         // true
		"a",        // true
		"AbCd",     // true
	}

	for _, test := range tests {
		fmt.Printf("%q -> %t\n", test, hasUniqueChars(test))
	}
}
