package wait

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServerはURLのサーバへ接続を試みます。
// 指数バックオフを使って一分間試みます。
// 全ての試みが失敗したらエラーを報告します。
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //成功
		}
		log.Printf("server not responding (%s);  retrying...", err)
		time.Sleep(time.Second << uint(tries)) //指数バックオフ
	}
	return fmt.Errorf("server %s failed to response after %s", url, timeout)
}

//main関数内で
func sample1() {
	url := ""
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}

func sample2() {
	url := ""
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down %v\n", err) //便利
		log.SetPrefix("wait: ")
		log.SetFlags(0)
	}
}

//エラーを記録しておく
func sample3() {
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v networking disabled", err)
	}

	//標準エラー
	if err := Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
	}
}

//エラーを無視する
func sample4() {
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir:%v", err)
	}

	//tempディレクトリをしよう
	os.RemoveAll(dir) //エラーを無視、$TMPDIRは定期的に削除される
}

func salmple5() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break //読み込み終了
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		//rを使う
		fmt.Print(r)
	}
}
