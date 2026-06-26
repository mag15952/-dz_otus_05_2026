package main

import (
	"fmt"
	"strings"
)

func main() {

	//var s string

	//s := []string{string('\u265a'), string('\u265b'), string('\u265c'),
	//	string('\u265d'), string('\u265e'), string('\u265f')}
	schess := []string{string('\u265c'), string('\u265e'), string('\u265d'),
		string('\u265a'), string('\u265b'), string('\u265d'), string('\u265e'), string('\u265c')}

	length := 15

	var result []string
	//pawn := string('\u265f')

	massiveCount := len(schess)
	count := 0

	for i := 0; i <= length; i++ {

		if count >= massiveCount {
			count = 0
		}

		result = append(result, schess[count])

		count++

	}

	fmt.Printf("%s\n", strings.Join(result, " | "))

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
