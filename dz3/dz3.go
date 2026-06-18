package main

import (
	"fmt"
	"strings"
)

func main() {
	var len int = 8 // нулевое значение
	//var y int // нулевое значение
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

		for j := 0; j < len; j++ {
			if strings.HasSuffix(str, "#") {
				str = str + " "
			} else {
				str = str + "#"
			}
		}

		fmt.Printf("%s\n", str)
	}

}
