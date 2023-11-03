package ex1678

// 40 ( 41 ) a 97 l 108 G 71 o 111

func Interpret(command string) string {

	commandRune := []rune(command)

	var result []rune

	for i := 0; i < len(commandRune); i++ {

		switch commandRune[i] {
		case rune(71):
			result = append(result, commandRune[i])
		case rune(40):
			if commandRune[i+1] == rune(41) {
				result = append(result, rune(111))
			} else {
				result = append(result, rune(97), rune(108))
			}
		default:

		}

	}

	return string(result)

}
