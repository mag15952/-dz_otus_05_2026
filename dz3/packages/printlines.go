package printlines

import (
	"math/rand"
	"strings"
)

func Printline(len int, str string) string {

	for j := 0; j < len-1; j++ {

		str = str + "|"

		if strings.HasSuffix(str, "#|") {
			str = str + " "
		} else {
			str = str + "#"
		}
	}
	return str
}

func PrintWhite(length int) string {

	s := []string{string('\u2654'), string('\u2655'), string('\u2656'),
		string('\u2657'), string('\u2658'), string('\u2659')}

	randomstr := makeRandomStr(length, s)

	return randomstr
}

func makeRandomStr(length int, s []string) string {

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
}

func randomInt(length int) int {
	var randInt int
	randInt = rand.Intn(length)
	return randInt
}

func PrintBlack(length int) string {

	s := []string{string('\u265a'), string('\u265b'), string('\u265c'),
		string('\u265d'), string('\u265e'), string('\u265f')}

	randomstr := makeRandomStr(length, s)

	return randomstr
}

func PrintLetters(length int) string {

	var s string
	var massive []string

	for i := 'A'; i <= 'Z'; i++ {

		massive = append(massive, string(i))

		/*if count > length { break }

		symbol := string(i)

		s = checkContains(s, symbol)
		s = s + " "

		count ++*/

	}

	massiveCount := len(massive)
	count := 0

	for i := 0; i <= length; i++ {

		if count >= massiveCount {
			count = 0
		}

		symbol := massive[count]

		s = checkContains(s, symbol)
		s = s + " "

		count++

	}

	return s
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
