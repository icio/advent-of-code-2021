package main

import (
	"fmt"
	"io"
)

func main() {
	const w = 3 // w = 1 is equivalent to d1p1.
	var d [w]int
	var incrs int

	for i := 0; i < w; i++ {
		_, err := fmt.Scan(&d[i])
		if err != nil {
			panic(err)
		}
	}
	for {
		var nextDepth int
		_, err := fmt.Scan(&nextDepth)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if nextDepth > d[0] {
			incrs++
		}
		copy(d[0:], d[1:])
		d[w-1] = nextDepth
	}
	fmt.Println(incrs)
}
