package main

import (
	"fmt"
	"io"
)

func main() {
	var d [3]int
	var incrs int

	_, err := fmt.Scan(&d[0], &d[1], &d[2])
	if err != nil {
		panic(err)
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
		d[0], d[1], d[2] = d[1], d[2], nextDepth
	}
	fmt.Println(incrs)
}
