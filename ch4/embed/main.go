package main

import "fmt"

type Point struct {
	X, Y int
}

/*
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}

*/

//無名フィールド形式
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8      //w.Circle.Center.X = 8と同じ
	w.Y = 8      //w.Circle.Center.Y = 8と同じ
	w.Radius = 5 //w.Circle.Radius = 5と同じ
	w.Spokes = 20

	// w = Wheel{8, 8, 5, 20}                       //コンパイルエラー
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} //コンパイルエラー
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, //カンマが必要
	}
	fmt.Printf("%#v\n", w) //#はアドヴァーブ、Go構文とにた形式で値を表維持させている
	w.X = 42
	fmt.Printf("%#v\n", w)
}
