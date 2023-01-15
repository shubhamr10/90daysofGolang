package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// modify dup2 to print all the filenames in which each duplicate line occurs.
func main() {
	var counts = map[string]map[string]int{}
	files := os.Args[1:]
	for _, filename := range files {
		counts[filename] = map[string]int{}
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex4:%v\n", err)
			continue
		}
		countLinesWithFileName(f, filename, counts)
		f.Close()
	}
	for line, n := range counts {
		log.Println(line, n)
	}
}

func countLinesWithFileName(f *os.File, name string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}
