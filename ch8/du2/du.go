package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose prgress messages")

func main() {
	//最初のディレクトリを決める。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//ファイルツリーを走査する
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()

	//定期的に結果を表示する
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
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
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents はディレクトリdirの項目を返します。
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1, %v\n", err)
		return nil
	}
	return entries
}
