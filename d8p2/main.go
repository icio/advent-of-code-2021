package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var sum int64
	for {
		// Scan the input.
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
			panic(fmt.Errorf("failed to scan 15 values"))
		}

		// Pick out the obvious digits.
		ds := [10]digit{
			8: wordnum("abcdefg"),
		}
		for _, word := range sig {
			d := wordnum(word)
			switch len(word) {
			// case 7: ds[8] = d // Always "abcdefg" in some order.
			case 2:
				ds[1] = d
			case 3:
				ds[7] = d
			case 4:
				ds[4] = d
			}
		}

		/*
			  0:        1:        2:        3:        4:
			 aaaa      ....      aaaa      aaaa      ....
			b    c    .    c    .    c    .    c    b    c
			b    c    .    c    .    c    .    c    b    c
			 ....      ....      dddd      dddd      dddd
			e    f    .    f    e    .    .    f    .    f
			e    f    .    f    e    .    .    f    .    f
			 gggg      ....      gggg      gggg      ....

			  5:        6:        7:        8:        9:
			 aaaa      aaaa      aaaa      aaaa      aaaa
			b    .    b    .    .    c    b    c    b    c
			b    .    b    .    .    c    b    c    b    c
			 dddd      dddd      ....      dddd      dddd
			.    f    e    f    .    f    e    f    .    f
			.    f    e    f    .    f    e    f    .    f
			 gggg      gggg      ....      gggg      gggg

			Here we can see that "3" contains "1" because all of the segments
			used by "1" are also used by "3".
		*/

		// Use the obvious digits to figure out the other digits.
		var d2or5 digit
		for _, word := range sig {
			d := wordnum(word)
			switch len(word) {
			// "1", "4", "7", "8" were identified above.
			case 2, 3, 4, 7:
				continue

			// "2", "3" and "5" have 5 segments.
			case 5:
				// "3" contains "1" whereas "2" and "5" do not.
				if d&ds[1] == ds[1] {
					ds[3] = d
					continue
				}

				// We need to compare "2" with "5", so store for later.
				if d2or5 == 0 {
					d2or5 = d
					continue
				}

				// "5" has more overlapping segments with "4" than "2" has.
				if bits(d&ds[4]) > bits(d2or5&ds[4]) {
					ds[5], ds[2] = d, d2or5
					continue
				}
				ds[2], ds[5] = d, d2or5

			// "0", "6" and "9" have 6 segments.
			case 6:
				// "6" does not contain "1" whereas "0" and "9" do.
				if d&ds[1] != ds[1] {
					ds[6] = d
					continue
				}
				// "9" contains "4" whereas "0" does not.
				if d&ds[4] == ds[4] {
					ds[9] = d
					continue
				}
				ds[0] = d
			}
		}

		// Convert the out words into digits.
		outnum := 0
		base := 1000
		for _, word := range out {
			d := wordnum(word)
			for n, pd := range ds {
				if d == pd {
					outnum += base * n
					break
				}
			}
			base /= 10
		}
		fmt.Println(strings.Join(out[:], " "), "=", outnum)

		sum += int64(outnum)
	}

	fmt.Println(sum)
}

type digit uint8

func bits(d digit) (n int) {
	for d > 0 {
		n++
		d &= d - 1
	}
	return
}

func (d digit) String() string {
	var r [8]rune
	for i := 0; i < 7; i++ {
		if d&digit(1<<i) != 0 {
			r[i] = 'a' + rune(i)
		} else {
			r[i] = '_'
		}
	}
	return string(r[:])
}

func wordnum(word string) (num digit) {
	for _, c := range word {
		num |= 1 << (c - 'a')
	}
	return
}
