package ex1061

import "fmt"

// R 82 D 68 L 76 U 85

func Smollest(s1 string, s2 string, baseStr string) string {
	var str []byte
	var min byte
	min = 122
	alphavit := make(map[byte][]byte)
	for i := 0; i < len(s1); i++ {
		alphavit[s1[i]] = append(alphavit[s1[i]], s2[i])
		alphavit[s2[i]] = append(alphavit[s2[i]], s1[i])
	}
	for i := 0; i < len(baseStr); i++ {
		fmt.Println(i)
		str = append(str, find(baseStr[i], alphavit))

	}
	fmt.Println(alphavit, "\n", min)
	return string(str)
}

func find(letter byte, alphavit map[byte][]byte) byte {
	/* for key, value := range alphavit {
		if letter == key {
			for _, v := range value {
				fmt.Println("Значение", v, letter)
				if v < letter {
					min = find(v, alphavit)
				} else {
					return letter
				}
			}
		}
	} */
	for _, v := range alphavit[letter] {
		fmt.Println("Значение", v, letter)
		if letter == v {
			continue
		} else if letter > v {
			letter = find(v, alphavit)
		} else {
			return letter
		}
	}

	return letter
}

/* for _, v := range alphavit[letter] {
	fmt.Println("Значение", v, letter)

	if letter == v {
		continue
	} else if min > v {
		min = v
		find(v, alphavit, min)
	}
	if len(alphavit[letter]) == 1 {
		break
	}

} */
