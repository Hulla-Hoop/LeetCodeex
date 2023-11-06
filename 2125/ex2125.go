package ex2125

// 0 48 1 49

func NumberOfBeams(bank []string) int {

	var (
		out   int
		laser int
		box   int
	)

	for i := 0; i < len(bank); i++ {
		countBox := 0
		for g := 0; g < len(bank[i]); g++ {
			if bank[i][g] == 49 {
				countBox++
			}
		}
		if box > 0 {
			laser = countBox * box
		}
		if countBox > 0 {
			box = countBox
		}
		out = out + laser
	}

	return out
}
