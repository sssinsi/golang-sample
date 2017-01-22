package methodvalue

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func sample() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance //メソッド値
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy //メソッド値
	scaleP(2)
	scaleP(3)
	scaleP(10)
}
