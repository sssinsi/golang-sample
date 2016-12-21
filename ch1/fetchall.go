package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //ゴルーチン開始
	}

	fp, err := os.OpenFile("res2.txt", os.O_WRONLY|os.O_APPEND, 0644)
	defer fp.Close()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	w := bufio.NewWriter(fp)

	for range os.Args[1:] {
		// fmt.Println(<-ch) //チャネルから受信
		_, err := w.WriteString(<-ch)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

	}
	w.Flush()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //chチャネルへ送信
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
