package main

import (
	"fmt"
	"main/dz5/models/chessboard"
	players "main/dz5/models/players"
	"main/dz5/packages/common"
)

func main() {

	// инициализация
	var length int

	//  ввод значений переменных
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&length)

	err := common.CheckValue(length)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	chessboard := chessboard.NewChessboard(length)

	players := players.NewPlayers()
	players.SetPlayers()

	// формирование и вывод линий
	players.PrintPlayer(0)

	chessboard.PrintChessboard()

	players.PrintPlayer(1)

}
