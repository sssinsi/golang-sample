package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sssinsi/golang-sample/ch5/links"
)

func main() {
	//コマンドライン引数から開始して、
	//ウェブを幅優先でクロールする。
	breadthFirst(crawl, os.Args[1:])
}

//breadthFirstはworklistないのここの項目に対してfを呼び出します。
//fから返された全ての項目はworklistへ追加されます
//fは、それぞれの項目に対して高々一度しか呼び出されません。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
