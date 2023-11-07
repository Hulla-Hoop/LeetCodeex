package ex1347

import "fmt"

func MinSteps(s string, t string) int {

	f := make(map[rune]int)
	d := make(map[rune]int)

	for _, i := range s {
		f[i]++
	}
	for _, i := range t {
		d[i]++
	}

	out := 0

	fmt.Println(f, "\n", d)

	for i, o := range f {
		if o > d[i] {
			out = out + o - d[i]
		}

	}

	return out
}
