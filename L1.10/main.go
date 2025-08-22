package main

import (
	"fmt"
	"sort"
)

func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, t := range temps {
		key := int(t/10.0) * 10
		groups[key] = append(groups[key], t)
	}
	return groups
}

func printGroups(groups map[int][]float64) {
	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemperatures(temps)
	fmt.Println(groups)
	printGroups(groups)
}
