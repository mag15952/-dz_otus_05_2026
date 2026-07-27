package main

import (
	"fmt"
	"main/dz7/internal/models"
	"main/dz7/internal/repository"
	"main/dz7/internal/service"
	"sync"
)

func main() {

	// инициализация
	var length int
	var boardcount int

	//  ввод значений переменных
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&length)

	fmt.Print("Введите количество досок: ")
	fmt.Scan(&boardcount)

	err := service.CheckValue(length)

	err1 := service.CheckValue(boardcount)

	if err != nil || err1 != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	chbrd := models.NewChessboard(length)
	plr := models.NewPlayers()

	repository.SetObject(plr)
	repository.SetObject(chbrd)

	service.ClearTerminal()

	for range boardcount {

		// формирование и вывод линий
		repository.PrintObject(plr, 0)
		repository.PrintObject(chbrd, 0)
		repository.PrintObject(plr, 1)
	}

	wg := sync.WaitGroup{}

	for range boardcount {
		defer wg.Done()
		go service.PlayChessNew(plr, length, chbrd)
	}

	wg.Wait()
}
