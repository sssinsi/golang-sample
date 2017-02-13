package bank2

var (
	sema    = make(chan struct{}, 1) //balanceを保護するバイナリセマフォ
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} //トークンを獲得
	balance = balance + amount
	<-sema //トークンを解放
}

func Balance() int {
	sema <- struct{}{} //トークンをッか疎く
	b := balance

	<-sema //トークンを解放
	return b
}
