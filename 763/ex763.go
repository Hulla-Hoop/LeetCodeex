package ex763

import "fmt"

// R 82 D 68 L 76 U 85

func PartitionLabels(s string) []int {
	var out []int
	end := len(s) - 1
	for i := 0; i < len(s[:end]); i++ {
		for g := end; g > 0; g-- {
			if s[i] == s[g] {
				end = g
				fmt.Println(end)
			}
		}

	}

	return out
}
