package methodexpression

import (
	"fmt"
	"math"
)

//Point is point
type Point struct{ X, Y float64 }

//Distance between p,q
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy is multiple
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   //メソッド式
	fmt.Println(distance(p, q))  //"5"
	fmt.Printf("%T\n", distance) //"func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale) //"func(*Point, float64)"
}
