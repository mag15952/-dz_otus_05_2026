package repository

import (
	"fmt"
)

type Player interface {
	PrintPlayer(int)
	SetPlayers(int)
}

type Chessbrd interface {
	PrintChessboard()
	SetChessboard()
}

type GameObject interface {
	SetObject(any, int)
	PrintObject(any, int)
}

func SetObject(obj any) {

	switch obj_type := obj.(type) {
	case Chessbrd:
		obj_type.SetChessboard()
	case Player:
		obj_type.SetPlayers(2)
	case nil:
		fmt.Println("Получен nil")
	default:
		fmt.Printf("Неизвестный тип: %T (значение: %v)\n", obj_type, obj_type)
	}
}

func PrintObject(obj any, pnum int) {

	switch obj_type := obj.(type) {
	case Chessbrd:
		obj_type.PrintChessboard()
	case Player:
		obj_type.PrintPlayer(pnum)
	case nil:
		fmt.Println("Получен nil")
	default:
		fmt.Printf("Неизвестный тип: %T (значение: %v)\n", obj_type, obj)
	}
}
