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
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			//何もしない
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
