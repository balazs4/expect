package main

import (
	"bufio"
	"fmt"
	"os"
)

type result struct {
	passed int
	failed int
	total  int
}

func main() {
	result := result{passed: 0, failed: 0, total: 0}
	stdin := bufio.NewScanner(os.Stdin)
	index := 0
	for stdin.Scan() {
		result.total++
		actual := stdin.Text()
		var expected string
		if len(os.Args) > index+1 {
			expected = os.Args[1+index]
		} else {
			expected = "<nil>"
		}
		index++
		if actual != expected {
			fmt.Fprintf(os.Stderr, "expected: '%s'\n  actual: '%s'\n\n", expected, actual)
			result.failed++
			continue
		}
		result.passed++
	}
	if result.total == 0 {
		fmt.Fprintln(os.Stderr, "no expectation, or no input")
		os.Exit(128)
	}
	os.Exit(result.failed)
}
