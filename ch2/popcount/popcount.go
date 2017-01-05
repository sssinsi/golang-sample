package popcount

//pc[i]はiポピュレーションカウントです
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}

//PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します。
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]+
    pc[byte(x>>(2*8))]+
    pc[byte(x>>(3*8))]+
    pc[byte(x>>(4*8))]+
    pc[byte(x>>(5*8))]+
    pc[byte(x>>(6*8))]+
    pc[byte(x>>(7*8))]+
    pc[byte(x>>(0*8))]+
    pc[byte(x>>(1*8))]
    )
}
