package main

import "fmt"

//構造体の全フィールドが比較可能ならば、構造体自体も比較可能
type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

func main() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	//比較可能な構造体ならば、mapのキーにも使えます
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
