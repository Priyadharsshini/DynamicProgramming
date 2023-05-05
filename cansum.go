package main

import "fmt"

func canSum(targetSum int, arr []int, hm *map[int]bool) bool {
	if val, ok := (*hm)[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}
	for _, valu := range arr {
		if canSum(targetSum-valu, arr, hm) == true {
			return true
		} else {
			(*hm)[targetSum] = false
		}
	}
	return false
}

func main() {
	hm := make(map[int]bool)
	fmt.Println(canSum(7, []int{5, 4, 3, 7}, &hm))
	fmt.Println(canSum(7, []int{2, 6}, &hm))
	fmt.Println(canSum(100, []int{3}, &hm))

}
