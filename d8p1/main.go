package main

import (
	"fmt"
	"io"
)

func main() {
	var easy int64

	for {
		var sig [10]string
		var sep string
		var out [4]string
		n, err := fmt.Scan(
			&sig[0], &sig[1], &sig[2], &sig[3], &sig[4], &sig[5], &sig[6], &sig[7], &sig[8], &sig[9],
			&sep,
			&out[0], &out[1], &out[2], &out[3],
		)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if n != 15 {
			panic(fmt.Errorf("failed to scan 10 values"))
		}

		for _, v := range out {
			switch len(v) {
			case 2, 4, 3, 7:
				easy++
			}
		}
	}

	fmt.Println(easy)
}
