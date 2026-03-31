package main

import (
	"flag"
	"fmt"
	"strings"
)

func echo(args []string, sep string, newline bool) string {
	out := strings.Join(args, sep)
	if newline {
		return out + "\n"
	}
	return out
}

func main() {
	sep := flag.String("s", " ", "separator")
	addNewline := flag.Bool("n", false, "add newline at end")
	noNewline := flag.Bool("N", false, "suppress newline at end")

	flag.Parse()

	newline := *addNewline
	if *noNewline {
		newline = false
	}

	result := echo(flag.Args(), *sep, newline)
	fmt.Print(result)
}
