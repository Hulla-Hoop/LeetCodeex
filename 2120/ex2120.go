package ex2120

// R 82 D 68 L 76 U 85

func dfs(n int, row int, column int, f []rune, countr int) int {
	if len(f) == 0 {
		return countr
	}
	switch f[0] {
	case 76: //L

		p := row - 1
		if p < 0 {
			return countr
		}
		countr++
		row = p
		countr = dfs(n, row, column, f[1:], countr)

	case 82: //R
		p := row + 1

		if p > n {
			return countr
		}
		countr++
		row = p
		countr = dfs(n, row, column, f[1:], countr)

	case 85: // U
		p := column - 1

		if p < 0 {
			return countr
		}
		countr++
		column = p
		countr = dfs(n, row, column, f[1:], countr)
	case 68: // D
		p := column + 1

		if p > n {
			return countr
		}
		countr++
		column = p
		countr = dfs(n, row, column, f[1:], countr)
	}

	return countr
}

func ExecuteInstructions(n int, startPos []int, s string) []int {
	row := startPos[1]
	column := startPos[0]
	var out []int
	for i, _ := range s {

		var countr int

		countr = dfs(n-1, row, column, []rune(s[i:]), countr)
		out = append(out, countr)
	}

	return out
}

// код с литкода
func ExecuteInstructionsTwo(n int, startPos []int, s string) []int {

	var res []int

	for i := 0; i < len(s); i++ {

		var count int

		// first way of copying slice
		//var pos = make([]int, len(startPos))
		//copy(pos, startPos)

		// second way of copying slice
		var pos = []int{startPos[0], startPos[1]}

		for ii := i; ii < len(s); ii++ {
			switch string(s[ii]) {
			case "U":
				pos[0]--
			case "L":
				pos[1]--
			case "R":
				pos[1]++
			case "D":
				pos[0]++
			}

			if pos[0] >= n || pos[0] < 0 || pos[1] >= n || pos[1] < 0 {
				break
			} else {
				count++
			}
		}

		res = append(res, count)
	}

	return res
}
