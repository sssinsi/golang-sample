package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	//最初のディレクトリを決める。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//ファイルツリーを走査する
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root = range roots {
		n.Add(1)
		go walkDir(root, &n, fileSize)
	}
	go func() {
		n.Wait()
		close(fileSize)
	}()
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop //fileSizeが閉じられた
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir はdirをルートとするファイルツリーをたどり、
// 見つかったファイルのそれぞれの大きさをfileSizesに送る
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//sema はdirentsでの平衡性を制限するための計数セマフォ
var sema = make(chan struct{}, 20)

// dirents はディレクトリdirの項目を返します。
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //tokenを獲得
	defer func() { <-sema }() //tokenを解放
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1, %v\n", err)
		return nil
	}
	return entries
}
