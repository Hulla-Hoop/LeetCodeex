package main

import "fmt"

//46 это точка 91-  [    93-  ]

func finalValueAfterOperations(operations []string) int {
	var summ int

	one := "++X"
	alterOne := "X++"
	minusOne := "--X"
	alterMinusOne := "X--"

	for _, o := range operations {
		if o == one || o == alterOne {
			summ += 1
		} else if o == minusOne || o == alterMinusOne {
			summ -= 1
		} else {
			fmt.Println("vex")
		}
	}

	return summ
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"X++", "X--"}))
}
