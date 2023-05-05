package main

import "fmt"

func gridTraveller(m int, n int, hm *map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)

	if val, ok := (*hm)[key]; ok {
		return val
	}
	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	(*hm)[key] = gridTraveller(m-1, n, hm) + gridTraveller(m, n-1, hm)
	return (*hm)[key]

}

func main() {
	hm := make(map[string]int)
	fmt.Println(gridTraveller(2, 3, &hm))
	fmt.Println(gridTraveller(4, 5, &hm))
	fmt.Println(gridTraveller(18, 32, &hm))

}
