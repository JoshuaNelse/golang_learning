// Dup prints the text of each line that appears more than
// once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func scanInputFromSTDIn() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoreing potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func scanInputFromSTDInOrFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// loop over files
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: Ignoring potential errors from input.Err()
}

func scanInputFromIOUtils() {
	counts := map[string]map[string]int{}
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		counts[filename] = map[string]int{}
		for _, line := range strings.Split(string(data), "\n") {
			counts[filename][line]++
		}
	}
	fmt.Printf("COUNT\tFILENAME\tDUPLICATE LINE TEXT\n")
	for filename, lines := range counts {
		for line, n := range lines {
			fmt.Printf("%d\t%s\t%s\n", n, filename, line)
		}
	}
}

// exercise 1 -- modify scan method to print all files in which duplicate lines occur
func exercise1() {
	scanInputFromIOUtils()
}

func main() {
	exercise1()
}
