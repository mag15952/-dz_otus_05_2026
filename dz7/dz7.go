package main

import (
	"fmt"
	"main/dz7/internal/models"
	"main/dz7/internal/repository2"
	"main/dz7/internal/service"
	common "main/dz7/internal/service/common"
	"sync"
	"time"
	//"time"
)

type GamePool struct {
	gamesChan chan *service.GameService
	wg        sync.WaitGroup
	stopOnce  sync.Once
	stopGame  chan struct{}
}

func main() {

	// инициализация
	var length int
	var gameCount int

	//  ввод значений переменных
	fmt.Print("Введите размер игрового поля: ")
	fmt.Scan(&length)

	err := common.CheckValue(length)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	fmt.Print("Введите количество игровых полей: ")
	fmt.Scan(&gameCount)

	err = common.CheckValue(gameCount)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	repo := repository2.NewRepo()
	players := models.NewPlayers("", "")
	players.SetPlayers(2)

	game := service.NewGame(repo)
	game.SetGame(length, gameCount, players.GetPlayer(0), players.GetPlayer(1))

	// каналы

	gp := &GamePool{
		gamesChan: make(chan *service.GameService),
		stopGame:  make(chan struct{}),
	}

	gp.wg.Add(gameCount)
	for i := 0; i < gameCount; i++ {

		go func() {
			service.PlayChess(gp.gamesChan, gp.stopGame, game)
		}()

	}

	for {
		select {
		case game, ok := <-gp.gamesChan:
			if !ok {
				return
			}
			defer gp.wg.Done()
			printTerminal(game)
		case <-gp.stopGame:
			return
		}
	}

	////////////////////////////////
	/*dataChan := make(chan *service.GameService)

	// Создаем тикер на 1 секунду
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Обязательно останавливаем при выходе

	var currentData *service.GameService

	// Горутина-эмулятор, которая шлет данные раз в 3 секунды
	go func() {
		//for i := 1; i <= 3; i++ {
		//	time.Sleep(3 * time.Second)
		//	dataChan <- game
		service.PlayChess(dataChan, game)
		//}
	}()

	// Главный цикл отрисовки и обработки событий
	for {
		select {
		case game := <-dataChan:
			// Обновляем текущее состояние при получении новых данных
			currentData = game

		case <-ticker.C:
			// Каждую секунду перерисовываем экран
			if currentData != nil {
				printTerminal(currentData)
			}
		}
	}
	*/

}

func printTerminal(game *service.GameService) {
	// Очистка экрана и возврат курсора в начало (ANSI-последовательность)
	//	fmt.Print("\033[H\033[2J")

	game.PrintAll()

	fmt.Println("Время:", time.Now().Format("15:04:05"))
}
