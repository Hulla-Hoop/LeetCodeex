package ex2114

func MostWordsFound(sentences []string) int {

	var maxCount int32
	for i := 0; i < len(sentences); i++ {
		var counter int32
		tempSl := []rune(sentences[i])
		for _, o := range tempSl {
			if o == rune(32) {
				counter++
			}
		}
		if counter > maxCount {
			maxCount = counter
		}
	}
	return int(maxCount + 1)
}
