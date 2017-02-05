package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//abortチャネルの作成
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //1バイトを読み込む
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	//なにもしない
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()

}

func launch() {
	fmt.Println("Lift off!")
}
