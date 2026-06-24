package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string

	/*for i := 'A'; i <= 'Z'; i++ {
		s = s + string(i)
		fmt.Printf("%c ", i)
	}
	//fmt.Println()*/

	for i := 'A'; i <= 'Z'; i++ {

		symbol := string(i)

		s = checkContains(s, symbol)
		s = s + " "

	}
	fmt.Printf("%s ", s)
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
