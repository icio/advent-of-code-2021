package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var errscore int
	stack := make([]byte, 0, 1000)

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
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
					errscore += charscore(c)
					break Chars
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	if err := stdin.Err(); err != nil {
		die(fmt.Errorf("reading stdin: %w", err))
	}
	fmt.Println(errscore)
}

func die(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func charscore(c byte) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
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
