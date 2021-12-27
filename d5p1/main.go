package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var h, w int

	var lines []line
	for {
		var l line
		n, err := fmt.Scanf("%d,%d -> %d,%d\n", &l.x1, &l.y1, &l.x2, &l.y2)
		if err == io.EOF {
			if n > 0 {
				panic(fmt.Errorf("read partial line"))
			}
			break
		}
		if err != nil {
			panic(err)
		}

		if l.x1 >= w {
			w = l.x1 + 1
		}
		if l.x2 >= w {
			w = l.x2 + 1
		}
		if l.y1 >= h {
			h = l.y1 + 1
		}
		if l.y2 >= h {
			h = l.y2 + 1
		}

		if l.x1 != l.x2 && l.y1 != l.y2 {
			fmt.Fprintf(os.Stderr, "Skipping non-vertical/horizontal line: %s\n", l)
			continue
		} else {
			// Sort lines - only works if vertical or horizontal.
			if l.x2 < l.x1 {
				l.x1, l.x2 = l.x2, l.x1
			}
			if l.y2 < l.y1 {
				l.y1, l.y2 = l.y2, l.y1
			}
		}

		lines = append(lines, l)
	}

	g := newGrid(h, w)
	for _, l := range lines {
		if l.x1 != l.x2 {
			row := g[l.y1]
			for x := l.x1; x <= l.x2; x++ {
				row[x]++
			}
		}
		if l.y1 != l.y2 {
			for y := l.y1; y <= l.y2; y++ {
				g[y][l.x1]++
			}
		}
	}

	n := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if g[y][x] >= 2 {
				n++
			}
		}
	}
	fmt.Println(n)
}

type grid [][]int

func newGrid(h, w int) grid {
	g := make(grid, h)
	for y := 0; y < h; y++ {
		g[y] = make([]int, w)
	}
	return g
}

type line struct{ x1, y1, x2, y2 int }

func (l line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.x1, l.y1, l.x2, l.y2)
}
