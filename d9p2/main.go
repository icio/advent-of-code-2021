package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	bitMark  = 0b1000000
	maskElev = 0b0111111
)

type grid struct {
	w int
	m []uint8
}

func (g grid) xy(x, y int) int {
	return y*g.w + x
}

func (g grid) elev(x, y int) uint8 {
	return g.m[g.xy(x, y)] & maskElev
}

func (g grid) mark(x, y int) {
	g.m[g.xy(x, y)] |= bitMark
}

func (g grid) marked(x, y int) bool {
	return g.m[g.xy(x, y)]&bitMark != 0
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
		m: make([]uint8, h*w),
	}
	n := 0
	for _, c := range stdin {
		if c < '0' || c > '9' {
			continue
		}
		g.m[n] = uint8(c - '0')
		n++
	}
	if n != h*w {
		die(fmt.Errorf("expected grid %d x %d but read only %d heights", h, w, n))
	}

	type basin struct{ x, y, size int }
	var bs [3]basin

	type coord struct{ x, y int }
	q := make([]coord, 0, h*w)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Check if we're at a low point.
			e := g.elev(x, y)
			if x > 0 && g.elev(x-1, y) <= e {
				continue
			}
			if x < w-1 && g.elev(x+1, y) <= e {
				continue
			}
			if y > 0 && g.elev(x, y-1) <= e {
				continue
			}
			if y < h-1 && g.elev(x, y+1) <= e {
				continue
			}

			// We're at the lowest point of a basin. Fan out to determine its
			// size.
			b := basin{x: x, y: y}
			g.mark(x, y)
			q = append(q, coord{x, y})
			for len(q) > 0 {
				c := q[len(q)-1]
				q = q[:len(q)-1]
				b.size++

				if c.x > 0 && !g.marked(c.x-1, c.y) && g.elev(c.x-1, c.y) < 9 {
					g.mark(c.x-1, c.y)
					q = append(q, coord{c.x - 1, c.y})
				}
				if c.x < w-1 && !g.marked(c.x+1, c.y) && g.elev(c.x+1, c.y) < 9 {
					g.mark(c.x+1, c.y)
					q = append(q, coord{c.x + 1, c.y})
				}
				if c.y > 0 && !g.marked(c.x, c.y-1) && g.elev(c.x, c.y-1) < 9 {
					g.mark(c.x, c.y-1)
					q = append(q, coord{c.x, c.y - 1})
				}
				if c.y < h-1 && !g.marked(c.x, c.y+1) && g.elev(c.x, c.y+1) < 9 {
					g.mark(c.x, c.y+1)
					q = append(q, coord{c.x, c.y + 1})
				}
			}

			fmt.Fprintf(os.Stderr, "(%d, %d) = %d => %d\n", x, y, e, b.size)

			// Keep the three largest basins.
			switch {
			case b.size > bs[0].size:
				bs[0], bs[1], bs[2] = b, bs[0], bs[1]
			case b.size > bs[1].size:
				bs[1], bs[2] = b, bs[1]
			case b.size > bs[2].size:
				bs[2] = b
			}
		}
	}

	fmt.Println(bs)
	fmt.Println(bs[0].size * bs[1].size * bs[2].size)
}

func die(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
