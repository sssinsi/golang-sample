package crawl1

import (
	"fmt"
	"log"

	"github.com/sssinsi/golang-sample/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
