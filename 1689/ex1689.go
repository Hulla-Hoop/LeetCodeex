package ex1689

func MinPartitions(n string) int {

	var g byte
	for i := 0; i < len(n); i++ {
		if n[i] == 57 {
			return 9
		}
		if n[i] > g {
			g = n[i]
		}
	}

	return int(g - '0')
}
