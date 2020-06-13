//Count duplicate lines either in files or stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	printed := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for name, linecount := range counts {
		for line, n := range linecount {
			if n > 1 {
				if printed[name] == false {
					fmt.Printf("Name of the file with duplicates: %v\n", name)
					printed[name] = true
				}
				fmt.Printf("%v\t%v\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	counts[f.Name()] = make(map[string]int)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}
