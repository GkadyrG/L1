package main

import "fmt"

func main() {
	nums := []int{3, 1, -11, 12, 12, 27, 1, -3, 4}
	sorted := quickSort(nums)

	fmt.Println(sorted)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]

	var left []int
	var right []int

	for _, num := range arr[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	result := append(sortedLeft, pivot)
	result = append(result, sortedRight...)

	return result
}
