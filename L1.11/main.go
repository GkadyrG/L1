package main

import "fmt"

func intersection(in1, in2 []int) []int {
	res := make([]int, 0)
	mp := make(map[int]struct{})

	for _, val := range in1 {
		mp[val] = struct{}{}
	}

	for _, val := range in2 {
		if _, ok := mp[val]; ok {
			res = append(res, val)
		}
	}

	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3}, []int{2, 3, 4}))
}
