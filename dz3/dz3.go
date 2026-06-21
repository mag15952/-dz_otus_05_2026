package main

import (
	"fmt"
	"strings"
)

func main() {

	var len int

	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&len)

	if len == 0 {
		len = 8
	}

	var str string

	countflag := true

	for i := 0; i < len; i++ {

		if countflag {
			str = " "
			countflag = false
		} else {

			str = "#"
			countflag = true
		}

		str = printline(len, str)

		fmt.Printf("%s\n", str)
	}

}

func printline(len int, str string) string {

	for j := 0; j < len-1; j++ {
		if strings.HasSuffix(str, "#") {
			str = str + " "
		} else {
			str = str + "#"
		}
	}
	return str
}
