package main

import "fmt"

//46 это точка 91-  [    93-  ]

func numJewelsInStones(jewels string, stones string) int {
	var summ int
	for i := 0; i < len([]rune(stones)); i++ {
		for u := 0; u < len([]rune(jewels)); u++ {
			if []rune(stones)[i] == []rune(jewels)[u] {
				summ += 1
			}
		}
	}
	return summ
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAaAaxcc"))
}
