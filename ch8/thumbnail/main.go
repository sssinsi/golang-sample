// +build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/sssinsi/golang-sample/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin) //set images path file
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text()) //scan row
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

// makeThumbnails は指定されたファイルのサムネイルを作成します。
// embarrassingly parallel
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//注意；正しくない！
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		//全ての処理が完了する前に処理が戻る
		go thumbnail.ImageFile(f) //注意：エラーを無視
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		//単一変数fは全ての無名関数値で共有されている。
		//連続したループの繰り返しで更新される。
		//(f)で明示的なパラメータを追加することで、go文が実行された時のfの値を使うことを保証しています。
		go func(f string) {
			thumbnail.ImageFile(f) //エラーを無視
			ch <- struct{}{}
		}(f)
	}
	//ゴルーチンの完了を待つ
	for range filenames {
		<-ch
	}
}

// makeThumbnails4は指定されたファイルのサムネイルを並列に作成します。
// なんらかの処理が失敗したらエラーを返します。
func makeThumbnails4(filenames []string) {
	//エラーチャネル
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			//最初のnilでないエラーを見つけると、そのエラーを呼び出し元へ返す。
			//(呼び出し元には)errorsチャネルを空にするゴルーチンはいない。
			//残ったワーカーのゴルーチンは、そのチャネルへ値を送ろうとして永久に待たされ、終了することはありません。
			//このゴルーチンのリークはプログラム全体を止めてしまったり、モメリを枯渇させたいるするかもしれません。
			//簡単な解決方法は、わかーのゴルーチンがメッセージを送る際に待たされることがないように、
			//十分な大きさの容量を持つバッファありチャネルを使うことです。
			//別の方法は、メインゴルーチンが最初のエラーを遅延なく返す一方で、チャネルを空にするための別のゴルーチンを生成することです。
			return err //正しくない：ゴルーチンのリーク！

		}
	}
	return nil
}

//makeThumbnails5 は指定されたファイルのサムネイルを並列に作ります。
//任意の順に作成されたファイル名か、または処理が失敗した場合はエラーを返します。
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

//makeThumbnails6 は指定されたファイルのサムネイルを並列に作ります。
//生成したファイルのバイト数を返します。
//カウンタ型 sync.WaitGroup を使います。
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup //活動中のゴルーチンの数
	for f := range filenames {
		//カウンタをあげる
		//ワーカーのゴルーチンのの外で呼びださなければならない。
		//そうしないとaAddの呼び出しがクローザのゴルーチンによるWaitの呼び出しより前に発生することが保証されない。
		wg.Add(1)
		//ワーカー(worker)
		go func(f string) {
			defer wg.Done() //カウンタを下げる
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) //ファイル情報取得。エラーを無視
			sizes <- info.Size()
		}(f)
	}
	//クローザー(closer)
	//sizesチャネルを閉じる前に、ワーカーの終了を待つクローザのゴルーチンを生成
	go func() {

		wg.Wait() //カウンタが0になるまで待つ
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
