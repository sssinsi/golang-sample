package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("aaa")
	counts := make(map[string]int)
	fmt.Println("aaa")
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("aaa")
	for input.Scan() {
		fmt.Println("aaa")
		counts[input.Text()]++
	}

	//input.Err()からのエラーの可能性をむししている
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
