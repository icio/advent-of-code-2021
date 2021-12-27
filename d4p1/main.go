package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	h = 5
	w = 5

	drawnBit = uint8(1 << 7)
	nMask    = ^drawnBit
)

type board [h][w]uint8

func (b board) find(n uint8) (ok bool, x, y int) {
	for y := 0; y < h; y++ {
		for x := x; x < w; x++ {
			if b[y][x] == n {
				return true, x, y
			}
		}
	}
	return false, 0, 0
}

func (b *board) drawCoord(x, y int) {
	b[y][x] |= drawnBit
}

func (b board) hasLine() bool {
Rows:
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if _, drawn := b.read(x, y); !drawn {
				continue Rows
			}
		}
		return true
	}

Cols:
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if _, drawn := b.read(x, y); !drawn {
				continue Cols
			}
		}
		return true
	}

	return false
}

func (b board) read(x, y int) (n uint8, drawn bool) {
	v := b[y][x]
	return v & nMask, v&drawnBit != 0
}

func (b board) sumUndrawn() (sum uint64) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			n, drawn := b.read(x, y)
			if drawn {
				continue
			}
			sum += uint64(n)
		}
	}
	return sum
}

func (b board) prettyPrint(f io.Writer, hx, hy int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			n, drawn := b.read(x, y)
			fmt.Fprintf(f, "% 3d", n)
			if drawn {
				if x == hx && y == hy {
					io.WriteString(f, "# ")
				} else {
					io.WriteString(f, "* ")
				}
			} else if x < w-1 {
				io.WriteString(f, "  ")
			}
		}
		fmt.Fprintln(f)
	}
}

func main() {
	// Read which numbers are drawn.
	var drawLine string
	_, err := fmt.Scan(&drawLine)
	if err != nil && err != io.EOF {
		panic(err)
	}

	// And parse into numbers.
	draws := make([]uint8, strings.Count(drawLine, ",")+1)
	for i, d := range strings.Split(drawLine, ",") {
		u, err := strconv.ParseUint(d, 10, 8)
		if err != nil {
			panic(err)
		}
		draws[i] = uint8(u)
	}

	// Read the boards.
	var boards []board
	for {
		var b board
		n, err := fmt.Scan(
			&b[0][0], &b[0][1], &b[0][2], &b[0][3], &b[0][4],
			&b[1][0], &b[1][1], &b[1][2], &b[1][3], &b[1][4],
			&b[2][0], &b[2][1], &b[2][2], &b[2][3], &b[2][4],
			&b[3][0], &b[3][1], &b[3][2], &b[3][3], &b[3][4],
			&b[4][0], &b[4][1], &b[4][2], &b[4][3], &b[4][4],
		)
		if err == io.EOF {
			if n > 0 {
				panic(fmt.Errorf("premature EOF reading board"))
			}
			break
		}
		if err != nil {
			panic(err)
		}
		boards = append(boards, b)
	}

	// Find the board which is completed first.
	for i, d := range draws {
		fmt.Fprintln(os.Stderr, "Draw", d)
		for b := range boards {
			found, x, y := boards[b].find(d)
			if !found {
				continue
			}

			fmt.Fprintf(os.Stderr, "# board %d finds %d at (%d, %d):\n", b, d, x, y)
			boards[b].drawCoord(x, y)
			boards[b].prettyPrint(os.Stderr, x, y)

			if i < h || i < w {
				continue
			}
			if !boards[b].hasLine() {
				continue
			}

			fmt.Printf("Drawing %d on turn %d/%d completes board %d:\n", d, i+1, len(draws), b)
			boards[b].prettyPrint(os.Stdout, -1, -1)
			undrawn := boards[b].sumUndrawn()
			fmt.Printf("(Sum undrawn: %d) * (Last drawn: %d) = %d\n", undrawn, d, undrawn*uint64(d))
			return
		}
	}
}
