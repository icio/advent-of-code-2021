package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var scores []int
	stack := make([]byte, 0, 1000)

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		var err bool
		stack = stack[:0]
	Chars:
		for _, c := range stdin.Bytes() {
			switch c {
			default:
				die(fmt.Errorf("unexpected char: %q", c))
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')', ']', '}', '>':
				if len(stack) == 0 || stack[len(stack)-1] != opening(c) {
					err = true
					break Chars
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if err {
			// The line is corrupted, ignore it.
			continue
		}
		if len(stack) == 0 {
			// The line is complete, no need to consider its autoclose score.
			continue
		}

		// Calculate the autoclose score.
		var s int
		for i := len(stack) - 1; i >= 0; i-- {
			s = s*5 + charscore(closing(stack[i]))
		}
		scores = append(scores, s)
	}
	if err := stdin.Err(); err != nil {
		die(fmt.Errorf("reading stdin: %w", err))
	}

	// Find the median autoclose score.
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func die(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func charscore(c byte) int {
	switch c {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}

func closing(c byte) byte {
	switch c {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	return 0
}

func opening(c byte) byte {
	switch c {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	}
	return 0
}
