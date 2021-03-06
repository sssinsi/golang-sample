package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]int) //key=単語,value=ファイル名
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			names, sep := "", ""
			for name := range filenames[line] {
				names += sep + name
				sep = " "
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, names)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if filenames != nil {
			if filenames[t] == nil {
				filenames[t] = make(map[string]int)
			}
			filenames[t][f.Name()]++
		}

	}
}
