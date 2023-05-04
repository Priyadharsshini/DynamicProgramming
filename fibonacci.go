package main

import "fmt"

func fib(n int, hm *map[int]int) int {
	if val, ok := (*hm)[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}
	(*hm)[n] = fib(n-1, hm) + fib(n-2, hm)
	return (*hm)[n]
}

func main() {
	hm := make(map[int]int)
	fmt.Println(fib(9, &hm))
	fmt.Println(fib(50, &hm))

}
