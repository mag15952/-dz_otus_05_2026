package printlines

import (
	"strings"
)

func Printline(len int, str string) string {

	for j := 0; j < len; j++ {

		str = str + "|"

		if strings.HasSuffix(str, "#|") {
			str = str + " "
		} else {
			str = str + "#"
		}
	}

	result := strings.Split(str, "|")
	return strings.Join(result, " | ")
}

func PrintWhite(length int) string {

	schess := []string{string('\u2656'), string('\u2658'), string('\u2657'),
		string('\u2654'), string('\u2655'), string('\u2657'), string('\u2658'), string('\u2656')}

	newstr := PrintSlice(length, schess)

	return newstr
}

func PrintWhitePowns(length int) string {

	schess := string('\u2659')

	newstr := PrintPownSlice(length, schess)

	return newstr
}

func PrintBlackPowns(length int) string {

	schess := string('\u265f')

	newstr := PrintPownSlice(length, schess)

	return newstr
}

func PrintPownSlice(length int, schess string) string {

	var result []string

	for i := 0; i <= length; i++ {

		result = append(result, schess)

	}

	return strings.Join(result, " | ")
}

func PrintSlice(length int, schess []string) string {

	var result []string

	massiveCount := len(schess)
	count := 0

	for i := 0; i <= length; i++ {

		if count >= massiveCount {
			count = 0
		}

		result = append(result, schess[count])

		count++

	}

	return strings.Join(result, " | ")
}

/*func makeRandomStr(length int, s []string) string {

	str := ""

	var result []string

	for j := 0; j < length; j++ {

		randInt := randomInt(len(s))

		result = append(result, s[randInt])

		//str = str + s[randInt]
		//str = str + "|"
	}

	str = strings.Join(result, "|")

	return str
}*/

/*func randomInt(length int) int {
	var randInt int
	randInt = rand.Intn(length)
	return randInt
}*/

func PrintBlack(length int) string {

	schess := []string{string('\u265c'), string('\u265e'), string('\u265d'),
		string('\u265a'), string('\u265b'), string('\u265d'), string('\u265e'), string('\u265c')}

	newstr := PrintSlice(length, schess)

	return newstr

}

func PrintLetters(length int) string {

	var massive []string
	var result []string

	for i := 'A'; i <= 'Z'; i++ {

		massive = append(massive, string(i))

	}

	massiveCount := len(massive)
	count := 0

	for i := 0; i <= length; i++ {

		if count >= massiveCount {
			count = 0
		}

		symbol := massive[count]

		result = append(result, symbol)

		count++

	}

	return strings.Join(result, " | ")
}

func checkContains(s string, symbol string) string {

	cont := strings.Contains(s, symbol)
	if cont == true {

		symbol = symbol + symbol
		s = checkContains(s, symbol)

	} else {

		s = s + symbol
	}
	return s
}
