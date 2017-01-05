package main

import "fmt"

func main() {
	fmt.Println(f(15))
}

func f(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
