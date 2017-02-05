package crawl2

import (
	"fmt"
	"log"

	"github.com/sssinsi/golang-sample/ch5/links"
)

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
