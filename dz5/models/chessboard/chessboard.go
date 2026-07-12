package chessboard

import (
	"fmt"
	printlines "main/dz5/packages/printlines"
)

type Chessboard struct {
	length int
}

func NewChessboard(len int) Chessboard {
	return Chessboard{
		length: len,
	}
}

func (c *Chessboard) PrintChessboard() {

	var str string
	countflag := true

	letters := printlines.PrintLetters(c.length - 1)
	fmt.Printf("  %s\n", letters)

	for i := 0; i < c.length; i++ {

		if i == 0 {
			//  расставляем белые
			str = printlines.PrintWhite(c.length - 1)
			//  расставляем белые пешки
		} else if i == 1 {
			str = printlines.PrintWhitePowns(c.length - 1)

		} else if c.length > 4 && i > 1 && i < c.length-2 {
			//  рисуем поле
			str = printlines.PrintlineNew(c.length, &i)
			countflag = !countflag

		} else if i == c.length-2 {
			//  расставляем черные пешки
			str = printlines.PrintBlackPowns(c.length - 1)
		} else if i == c.length-1 {
			//  расставляем черные фигуры
			str = printlines.PrintBlack(c.length - 1)
		}

		fmt.Printf("%d %s\n", i+1, str)
	}
}
