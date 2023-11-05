package ex2391

import "fmt"

func GarbageCollection(garbage []string, travel []int) int {
	posMap := make(map[byte]int)
	countMap := make(map[byte]int)
	for i := 0; i < len(garbage); i++ {
		for g := 0; g < len(garbage[i]); g++ {
			posMap[garbage[i][g]] = i
			countMap[garbage[i][g]]++
		}

	}
	fmt.Println(posMap, countMap)
	var summ int
	for i, g := range posMap {
		for t := 0; t < g; t++ {
			summ = summ + travel[t]
		}
		summ = summ + countMap[i]
	}
	fmt.Println(summ)
	return summ
}
