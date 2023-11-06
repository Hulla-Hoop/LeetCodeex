package ex2405

// R 82 D 68 L 76 U 85

func PartitionString(s string) int {
	cheker := make(map[byte]int)
	counter := 0
	for i := 0; i < len(s); i++ {
		cheker[s[i]]++
		if cheker[s[i]] > counter {
			counter++
		} else {
			cheker[s[i]] = cheker[s[i-1]]
		}

	}
	return counter
}

func PartitionStringTwo(s string) int {
	cheker := make(map[byte]bool)
	counter := 1
	for i := 0; i < len(s); i++ {
		if cheker[s[i]] {
			counter++
			for f := range cheker {
				cheker[f] = false
			}
		}
		cheker[s[i]] = true

	}
	return counter
}

func PartitionStringThree(s string) int {

	return 2
}
