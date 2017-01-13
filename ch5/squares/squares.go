package main

import (
	"fmt"
)

//squaresは呼び出されるごとに次の平方根数を返す関数を返します

func squares() func() int {
	//xという状態をもつ。。。！！１
	//関数を参照型として分類し、関数値が比較可能ではない理由はこれらの隠蔽された変数への参照のため
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16
}
