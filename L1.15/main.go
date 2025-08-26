package main

import (
	"fmt"
	"runtime"
)

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createHugeString(10 << 20)

	justStringBytes := make([]byte, 100)
	copy(justStringBytes, v[:100])
	justString = string(justStringBytes)

}

func printMemUsage(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc = %v KB\n", tag, m.Alloc/1024)
}

func main() {
	printMemUsage("Before")
	someFunc()
	printMemUsage("After")
	fmt.Println("justString length:", len(justString))
}
