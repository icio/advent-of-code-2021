package main

import (
	"fmt"
	"io"
)

func main() {
	var depth int
	var incrs int

	_, err := fmt.Scan(&depth)
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
		if nextDepth > depth {
			incrs++
		}
		depth = nextDepth
	}
	fmt.Println(incrs)
}
