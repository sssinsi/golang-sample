package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sssinsi/golang-sample/ch5/links"
)

func main() {
	worklist := make(chan []string)  //URLのリスト、重複を含む
	unseenLinks := make(chan string) //重複してないURL

	//コマンドライン引数をworklistへ追加する
	go func() { worklist <- os.Args[1:] }()

	//未探索のリンクを取得するために20このクローラのゴルーチンを生成する
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	//メインゴルーチンはworklistの項目の重複をなくし、
	//未探索の項目をクローラへ送る。
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

//token は２０個の並行なリクエストという限界を矯正するために使われる計数セマフォ
//大きさがゼロのstruct{}とする
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //トークンを獲得
	list, err := links.Extract(url)
	<-tokens //トークンを解放
	if err != nil {
		log.Println(err)
	}
	return list
}
