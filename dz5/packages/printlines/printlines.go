package printlines

import (
	"main/dz5/packages/common"
	"strings"
)

func PrintlineNew(len int, countflag bool) string {

	var result []string

	if countflag {
		result = append(result, " ")
	} else {
		result = append(result, "#")
	}

	for j := 0; j < len; j++ {

		if result[j] == " " {
			result = append(result, "#")

		} else {
			result = append(result, " ")
		}
	}

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

		symbol := common.CheckContains(result, massive[count])

		result = append(result, symbol)

		count++

	}

	return strings.Join(result, " | ")
}
