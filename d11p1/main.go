package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		die(err)
	}

	// Read the digits from stdin.
	var g grid
	n := 0
	for _, c := range stdin {
		if c < '0' || c > '9' {
			continue
		}
		if n == 100 {
			die("got more than 100 digits")
		}
		g[n/w][n%w] = int(c - '0')
		n++
	}
	if n != 100 {
		die(fmt.Errorf("only got %d digits", n))
	}

	dc := [...]coord{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{0, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	// Perform the 100 steps.
	var flashes int64
	q := make([]coord, 0, 100)
	for s := 0; s < 100; s++ {
		if s%10 == 0 {
			fmt.Println("Step", s)
			g.WriteTo(os.Stdout)
		}

		// Increment all by 1.
		q = q[:0]
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				g[y][x]++
				if g[y][x] == 10 {
					q = append(q, coord{x, y})
				}
			}
		}

		// Increment the neighbours of those reaching 10.
		for len(q) > 0 {
			c := q[len(q)-1]
			q = q[:len(q)-1]

			for _, d := range dc {
				n := coord{c.x + d.x, c.y + d.y}
				if n.x < 0 || n.y < 0 || n.x >= w || n.y >= h {
					continue
				}
				g[n.y][n.x]++
				if g[n.y][n.x] == 10 {
					q = append(q, n)
				}
			}
		}

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if g[y][x] > 9 {
					flashes++
					g[y][x] = 0
				}
			}
		}
	}

	fmt.Println("Done")
	g.WriteTo(os.Stdout)
	fmt.Println(flashes, "flashes")
}

type coord struct{ x, y int }

type grid [10][10]int

const (
	h = len(grid{})
	w = len(grid{}[0])
)

func (g grid) WriteTo(f io.Writer) (n int64, err error) {
	v := make([]interface{}, w)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v[x] = g[y][x]
		}
		m, err := fmt.Fprintln(f, v...)
		n += int64(m)
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func die(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
