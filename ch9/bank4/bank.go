package bank4

import "sync"

var (
	mu      sync.RWMutex //balanceを保護する
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

//アトミックではない！
// func Withdraw(amount int) bool {
// 	Deposit(-amount)
// 	if Balance() < 0 {
// 		Deposit(amount)
// 		return false //残高不足
// 	}
// 	return true
// }

//正しくない！
// func Withdraw(amount int) bool {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	Deposit(-amount)
// 	if Balance() < 0 {
// 		Deposit(amount)//mu.Lockでデッドロックを起こす
// 		return false //残高不足
// 	}
// 	return true
// }

//この関数はロックが獲得されていることを前提としている。
func deposit(amount int) { balance += amount }

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false //残高不足
	}
	return true
}
