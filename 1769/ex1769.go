package ex1769

// 49 1 48 0

func MinOperations(boxes string) []int {
	t := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {

		if boxes[i] == 49 {
			for g := 0; g < len(t); g++ {
				f := g - i
				if f < 0 {
					f = f * -1
				}
				t[g] += f
			}
		}
	}

	return t
}

// 0 1 2
// 1 0 1
