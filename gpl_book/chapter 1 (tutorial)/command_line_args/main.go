// prints command line arguement like echo CLI command.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func buildStringWithForLoop(args []string) string {
	var b strings.Builder
	b.Grow(len(args))
	b.WriteString(args[0])
	for _, s := range args[1:] {
		b.WriteString(" ")
		b.WriteString(s)
	}
	return b.String()
}

func buildStringWithStringJoin(args []string) string {
	return strings.Join(args[1:], " ")
}

// exercise 1 -- print os.Args[0] as well
func exercise1() {
	fmt.Println("-- exercise 1 --")
	// os.Args[0] is the command itself
	// os.Args[1:] contains all the parameters
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

// exercise 2 -- print the index and value of each arguement (one per line)
func exercise2() {
	fmt.Println("-- exercise 2 --")
	for i, arg := range os.Args {
		fmt.Printf("[%d] %s\n", i, arg)
	}
}

// exercise 3 -- Measure difference between using string.Join and a for loop
func exercise3() {
	fmt.Println("-- exercise 3 --")

	// timing for loop to create string from an array
	startForLoop := time.Now()
	forLoopOutput := buildStringWithForLoop(os.Args)
	forLoopduration := time.Since(startForLoop)
	fmt.Printf("My code took [%v] to build the string \"%s\"\n", forLoopduration, forLoopOutput)

	// timing strings.Join to create string from array
	startStringJoin := time.Now()
	stringJoinOutput := buildStringWithStringJoin(os.Args)
	stringJoinduration := time.Since(startStringJoin)
	fmt.Printf("strings.Join took [%v] to build the string \"%s\"\n", stringJoinduration, stringJoinOutput)
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}
