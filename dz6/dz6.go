package main

import (
	"fmt"
	"main/dz6/internal/models/chessboard"
	"main/dz6/internal/models/players"
	common "main/dz6/internal/repository"
	"main/dz6/internal/service"
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

	player := players.NewPlayers()
	player.SetPlayers()

	common.ClearTerminal()

	// формирование и вывод линий
	player.PrintPlayer(0)

	chessboard.PrintChessboard()

	player.PrintPlayer(1)

	service.PlayChess(player, length, &chessboard)
}
