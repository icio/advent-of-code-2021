package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type grid struct {
	w int
	m []int
}

func (g grid) set(x, y, h int) {
	g.m[y*g.w+x] = h
}

func (g grid) e(x, y int) int {
	return g.m[y*g.w+x]
}

func main() {
	// Read stdin.
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		die(fmt.Errorf("reading stdin: %w", err))
	}

	// Calculate the height and width of the map.
	w := bytes.IndexRune(stdin, '\n')
	if w == -1 {
		die(fmt.Errorf("no line feed found on stdin"))
	}
	h := len(stdin) / (w + 1)

	// Read the heightmap from stdin.
	g := grid{
		w: w,
		m: make([]int, h*w),
	}
	n := 0
	for _, c := range stdin {
		if c < '0' || c > '9' {
			continue
		}
		g.m[n] = int(c - '0')
		n++
	}
	if n != h*w {
		die(fmt.Errorf("expected grid %d x %d but read only %d heights", h, w, n))
	}

	risk := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			e := g.e(x, y)
			if x > 0 && g.e(x-1, y) <= e {
				continue
			}
			if x < w-1 && g.e(x+1, y) <= e {
				continue
			}
			if y > 0 && g.e(x, y-1) <= e {
				continue
			}
			if y < h-1 && g.e(x, y+1) <= e {
				continue
			}
			fmt.Fprintf(os.Stderr, "(%d, %d) = %d\n", x, y, e)
			risk += e + 1
		}
	}
	fmt.Println(risk)
}

func die(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
