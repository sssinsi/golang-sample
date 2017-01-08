package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// s := []string{"ab", "c", "d"}
	// fmt.Println(fmt.Sprintf("%q", s))

	seen := make(map[string]bool) //文字列のセット
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dudup: %v\n", err)
		os.Exit(1)
	}
}

// var m = make(map[string]int)
// func k(list []string) string { return fmt.Sprintf("%q", list) }
// func Add(list []string){m[k(list)++]}
// func count(list []string)int{return m[k(list)]}
