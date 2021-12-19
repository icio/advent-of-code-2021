package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	// Read all of the binary values and convert into integers.
	bits := 0
	rs := make([]int32, 0, 1000)
	for {
		var line string
		_, err := fmt.Scan(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		bits = len(line)
		r, err := strconv.ParseInt(line, 2, 32)
		if err != nil {
			panic(err)
		}
		rs = append(rs, int32(r))
	}

	// Find g.
	var g int32
	opts := make([]int32, len(rs))
	copy(opts, rs)
	for pos := bits - 1; pos >= 0; pos-- {
		opts = filterBitPos(opts, pos, mostCommonBit(opts, pos))
		if len(opts) == 1 {
			g = opts[0]
			break
		}
		if len(opts) == 0 {
			panic("no opts remaining")
		}
		if pos == 0 {
			panic("too many opts remaining")
		}
	}
	fmt.Println("g = ", g)

	// Find s.
	var s int32
	opts = append(opts[:0], rs...)
	for pos := bits - 1; pos >= 0; pos-- {
		opts = filterBitPos(opts, pos, 1-mostCommonBit(opts, pos))
		if len(opts) == 1 {
			s = opts[0]
			break
		}
		if len(opts) == 0 {
			panic("no opts remaining")
		}
		if pos == 0 {
			panic("too many opts remaining")
		}
	}
	fmt.Println("s = ", s)

	// Life support rating.
	fmt.Println("life = ", g*s)
}

func mostCommonBit(set []int32, pos int) int32 {
	var n0, n1 int32

	m := int32(1 << pos)
	for _, n := range set {
		if n&m != 0 {
			n1++
		} else {
			n0++
		}
	}

	if n0 > n1 {
		return 0
	}
	return 1
}

func filterBitPos(set []int32, pos int, bit int32) []int32 {
	match := set[:0]
	m := int32(1 << pos)
	v := m * bit
	for _, n := range set {
		if n&m == v {
			match = append(match, n)
		}
	}
	return match
}
