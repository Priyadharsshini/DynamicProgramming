package main

import (
	"fmt"
)

func bestSum(targetSum int, arr []int, hm *map[int][]int) []int {
	if val, ok := (*hm)[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	shortestArray := []int(nil)

	for _, val := range arr {
		result := bestSum(targetSum-val, arr, hm)
		if result != nil {
			result = append(result, val)
			if shortestArray == nil || len(shortestArray) > len(result) {
				shortestArray = result
			}
		}
	}
	(*hm)[targetSum] = shortestArray
	return shortestArray
}

func main() {
	hm := make(map[int][]int)
	fmt.Println(bestSum(7, []int{2, 4, 5, 4, 3, 7}, &hm))
}
