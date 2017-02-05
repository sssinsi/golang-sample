// +build ignore

package main

import "os"

func main() {
	worklist := make(chan []string)
	var n int //worklistへの送信待ちの数

	//コマンドラインの引数で始める
	//カウンタnはまだ発生していないworklistへの送信数を記録する
	n++
	go func() { worklist <- os.Args[1:] }()
	//crawlのゴルーチンを開始する前にもまた加算する。
	//mainのループはnがゼロになった時にこれ以上は行う処理がないので終了する

	seen := make(map[string]bool)
	//webを並行にクロールする
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
