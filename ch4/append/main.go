package main

func main() {

}

func appendIng(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		//拡張する余地がある。スライスを拡張する
		z = x[:zlen]
	} else {
		// 十分な領域がない。新たな配列を割り当てる
		//計算量を線形に均するために倍に拡大する
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	//... 少なくともzlenまでzを拡張する
	copy(z[len(x):], y)
	return z

}
