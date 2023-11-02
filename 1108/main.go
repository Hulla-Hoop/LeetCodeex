package main

import "fmt"

//46 это точка 91-  [    93-  ]

func defangIPaddr(address string) string {

	runeAddress := []rune(address)

	for i := 0; i < len(runeAddress); i++ {

		if runeAddress[i] == rune(46) {
			temp := runeAddress[i+1:]
			temp2 := runeAddress[0:i]
			temp3 := defangIPaddr(string(temp))
			g := []rune{}
			g = append(g, temp2...)
			g = append(g, rune(91), runeAddress[i], rune(93))
			g = append(g, []rune(temp3)...)
			return string(g)
		}

	}
	return string(runeAddress)
}

func main() {
	fmt.Println(defangIPaddr("255.2.4.5"))
}
