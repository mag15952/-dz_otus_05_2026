package main

import (
	"fmt"
	"main/dz6/internal/models/chessboard"
	"main/dz6/internal/models/players"
	"main/dz6/internal/repository"
	"main/dz6/internal/service"
	common "main/dz6/internal/service/common"
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

	chbrd := chessboard.NewChessboard(length)
	plr := players.NewPlayers()

	repository.SetObject(plr)
	repository.SetObject(chbrd)

	common.ClearTerminal()

	// формирование и вывод линий
	repository.PrintObject(plr, 0)

	repository.PrintObject(chbrd, 0)

	repository.PrintObject(plr, 1)

	service.PlayChessNew(plr, length, chbrd)
}
