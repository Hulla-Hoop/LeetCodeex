package ex890

import (
	"fmt"
)

func FindAndReplacePattern(words []string, pattern string) []string {
	AMap := make(map[byte]byte)
	var out []string
	patMap := make(map[byte]byte)
	for i := 0; i < len(words); i++ {
		t := false
		fmt.Println(i)
		for g := 0; g < len(words[i]); g++ {
			if _, ok := AMap[pattern[g]]; !ok {
				AMap[pattern[g]] = words[i][g]
			}
			_, ok := patMap[words[i][g]]
			if !ok {
				patMap[words[i][g]] = pattern[g]
			}

		}

		if len(AMap) == len(patMap) {
			t = true
			for o, j := range AMap {
				fmt.Println("j--", o, "patMat---", patMap[j])
				if o != patMap[j] {
					t = false
				}
			}

		}
		if t == true {
			out = append(out, words[i])
		}
		t = true
		fmt.Println(AMap, patMap)
		for g := range AMap {
			delete(AMap, g)
		}
		for g := range patMap {
			delete(patMap, g)
		}

	}
	return out
}
