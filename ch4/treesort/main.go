package main

import "image/gif"
import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

type Point struct{ X, Y int }

type T struct{ a, b int }

func main() {
	p := Point{1, 2}
	anim := gif.GIF{LoopCount: nframes}

	//a,bは公開されてない
	//他のモジュールにて
	var _ = p.T{a: 1, b: 2} //コンパイルエラー
	var _ = p.T{1, 2}       //コンパイルエラー

	//構造体はポインタを通して扱われるのが普通。
	//struct変数を生成して初期化し、そしてそのアドレスを得るのには　次の短い表記を使うことができます
	//pp := &Point{1, 2}
    //これと同じことです
    pp: = new(Point)
    *pp = Point{1,2}
    //&Point{1,2}は関数呼び出しなどの式内で直接使うことができる

}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

//Sort Sortはvalues内の値をその中でソートします
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

//appendValuesはtの要素をvaluesの正しい順序に追加し、
//結果のスライスを返します。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		//return &tree{value: value}と同じ
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
