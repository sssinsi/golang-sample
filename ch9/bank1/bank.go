package main

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances } //送信されたら実行される

func teller() {
	var balance int //balance はtellerゴルーチンに閉じ込められる
	for {
		fmt.Println(balance)
		select {
		case amount := <-deposits: //送信されたら実行される
			balance += amount
		case balances <- balance: //受信の準備(Balanceが呼び出されたら)ができたら実行される
		}
	}
}

func main() {
	go teller() //モニターゴルーチンを開始する
	Deposit(100)
	Deposit(50)
	Balance()
	Balance()
	Balance()
	Balance()
	Balance()

}
