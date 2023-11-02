package ex1108

//46 это точка 91-  [    93-  ]

func DefangIPaddr(address string) string {

	runeAddress := []rune(address)

	for i := 0; i < len(runeAddress); i++ {

		if runeAddress[i] == rune(46) {
			temp := runeAddress[i+1:]
			temp2 := runeAddress[0:i]
			temp3 := DefangIPaddr(string(temp))
			g := []rune{}
			g = append(g, temp2...)
			g = append(g, rune(91), runeAddress[i], rune(93))
			g = append(g, []rune(temp3)...)
			return string(g)
		}

	}
	return string(runeAddress)
}
