package geometry

import "math"

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path は点を直線で結びつける道のりです。
type Path []Point

//Distance はpathに沿って進んだ距離を返します
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// SceleBy は倍にします
//ポインタレシーバ
//p *Pointはレシーバパラメータ
//慣習では PointPointnおどれかのmethodがポインタレシーバを持つのであれば、厳密には必要なくてもPointの全てのmethodはPointレシーバを持つべき。
//method名:(*Point).ScaleBy
func (p *Point) SceleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func sample() {
	r := &Point{1, 2} //アドレスを返して、レシーバrを作る、rは*Point型のポインタ変数
	r.SceleBy(2)

	//もしくは
	p := Point{1, 2}
	pptr := &p //pのアドレスを格納するpptrは*Point型のポインタ変数
	pptr.SceleBy(2)

	//あるいは
	p2 := Point{1, 2}
	(&p2).SceleBy(2)

	//レシーバpがPoint型の変数でも、methodが*Pointレシーバを要求する場合には、次の表記が使える
	//コンパイラが変数に対して暗黙的に&pを行う。これは変数に対してだけ行われる。
	p.SceleBy(2)
}
