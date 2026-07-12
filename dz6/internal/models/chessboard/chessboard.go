package chessboard

import (
	common "main/dz6/internal/repository"
	"strconv"
)

type Chessboard struct {
	length   int
	AllBoard [][]string
}

func NewChessboard(len int) Chessboard {
	return Chessboard{
		length: len,
	}
}

func (c *Chessboard) AddAllBoard(arr []string) {
	c.AllBoard = append(c.AllBoard, arr)
}

func (c *Chessboard) PrintChessboard() {

	getLetters(c) //  расставляем буквы

	for i := 0; i < c.length; i++ {

		if i == 0 {
			//  расставляем белые
			getWhite(c)
			//  расставляем белые пешки
		} else if i == 1 {
			getWhitePowns(c)

		} else if c.length > 4 && i > 1 && i < c.length-2 {
			//  рисуем поле

			getlineNew(c, &i)

		} else if i == c.length-2 {
			//  расставляем черные пешки
			getBlackPowns(c, c.length-1)

		} else if i == c.length-1 {
			//  расставляем черные фигуры
			getBlack(c, c.length)
		}

	}

	common.PrintSliceNew(c.AllBoard)

}

func getLetters(c *Chessboard) {

	var massive []string
	var result []string

	result = append(result, " ")

	for i := 'A'; i <= 'Z'; i++ {

		massive = append(massive, string(i))

	}

	massiveCount := len(massive)
	count := 0

	for i := 0; i <= c.length-1; i++ {

		if count >= massiveCount {
			count = 0
		}

		symbol := common.CheckContains(result, massive[count])

		result = append(result, symbol)

		count++

	}

	c.AddAllBoard(result)

}

func getlineNew(c *Chessboard, i *int) {

	var result []string
	num := strconv.Itoa(*i + 1)

	result = append(result, num)

	for j := 0; j < c.length; j++ {

		if (*i+j)%2 == 0 {
			result = append(result, "#")

		} else {
			result = append(result, " ")
		}
	}

	c.AddAllBoard(result)

}

func getWhite(c *Chessboard) {

	schess := []string{string('\u2656'), string('\u2658'), string('\u2657'),
		string('\u2654'), string('\u2655'), string('\u2657'), string('\u2658'), string('\u2656')}

	c.AddAllBoard(makeSlice(c, schess, 1))
}

func getWhitePowns(c *Chessboard) {

	schess := string('\u2659')

	c.AddAllBoard(getPownSlice(c.length, schess, 2))
}

func getBlackPowns(c *Chessboard, num int) {

	schess := string('\u265f')

	c.AddAllBoard(getPownSlice(c.length, schess, num))
}

func getBlack(c *Chessboard, num int) {

	schess := []string{string('\u265c'), string('\u265e'), string('\u265d'),
		string('\u265a'), string('\u265b'), string('\u265d'), string('\u265e'), string('\u265c')}

	c.AddAllBoard(makeSlice(c, schess, num))

}

func getPownSlice(length int, schess string, num int) []string {

	var result []string
	result = append(result, strconv.Itoa(num))

	for i := 0; i <= length-1; i++ {

		result = append(result, schess)

	}

	return result
}

func makeSlice(c *Chessboard, schess []string, num int) []string {

	var result []string

	result = append(result, strconv.Itoa(num))

	massiveCount := len(schess)
	count := 0

	for i := 0; i <= c.length-1; i++ {

		if count >= massiveCount {
			count = 0
		}

		result = append(result, schess[count])

		count++

	}

	return result
}
