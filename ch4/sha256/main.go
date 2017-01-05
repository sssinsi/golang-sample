package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	m()
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func m() {
	months := [...]string{
		1:  "January",
		2:  "Feburary",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Println(cap(months))
	fmt.Println(len(months))
	fmt.Printf("%T\n", months)
	fmt.Println(months[1:13])
	fmt.Println(months[1:])
	fmt.Println(months[:])
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:20])
	fmt.Println(cap(summer))
	fmt.Println(len(summer))
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
