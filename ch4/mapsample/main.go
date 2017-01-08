package main

import (
	"fmt"
	"sort"
)

func main() {
	sample()
}
func sample() {

	ages := map[string]int{
		"charlie": 34,
		"alice":   12,
	}

	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	if age, ok := ages["bob"]; !ok {
		//bobがkeyとして存在しなかった時
		fmt.Println("no bob")
	} else {
		fmt.Println(age)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
