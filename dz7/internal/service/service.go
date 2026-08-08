package service

import (
	"fmt"
	"log"
	"main/dz7/internal/models"
	"main/dz7/internal/repository2"
	common "main/dz7/internal/service/common"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

type GameService struct {
	repo repository2.Repo
	game *models.Game
}

func NewGame(repo repository2.Repo) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (g *GameService) SetGame(len, boardCount int, player1, player2 string) {

	g.game = models.NewGame(boardCount)
	players := models.NewPlayers(player1, player2)
	board := models.SetChessboard(len)
	g.game.SetGameBoard(board)
	g.game.SetGamePlayers(players)

	// Сохраняем в репозиторий
	g.repo.SaveObject(g.game, 0)
	g.repo.SaveObject(players, 0)
	g.repo.SaveObject(board, 0)

}

func PlayChess(dataChan chan *GameService, stop chan struct{}, g *GameService) {

	//chessboard := g.board
	board := g.game.GetGameBoard()
	allBoard := board.GetAllBoard()
	length := board.GetLen()
	players := g.game.GetGamePlayers()

	printBoardSlice(&players, allBoard)

	for {

		var newMove models.MoveStruct
		newMove = models.NewMove()

		for i := range 2 {

			p := players.GetPlayer(i)
			//cancel, movecount, move = setCommands(p)
			setCommands(p, &newMove, i)

			if newMove.Cancel == true {
				stop <- struct{}{}
				break
			} else if newMove.Movecount > 0 {

				m := map[string]int{
					"line_from":   0,
					"column_from": 0,
					"line_to":     0,
					"column_to":   0,
				}

				for movec := range newMove.Movecount {

					if movec%2 == 0 {

						switch i {
						case 0:
							{
								m["line_from"] = 1
							}

						case 1:
							{
								m["line_from"] = length
							}
						}
					} else {

						switch i {
						case 0:
							{
								m["line_from"] = 2
							}

						case 1:
							{
								m["line_from"] = length - 1
							}
						}
					}

					switch i {

					case 0:
						{
							min := 3
							max := length
							m["line_to"] = rand.IntN(max-min+1) + min
						}

					case 1:
						{

							min := 1
							max := length - 3
							m["line_to"] = rand.IntN(max-min+1) + min
						}
					}
					min := 1
					max := length
					m["column_from"] = rand.IntN(max-min+1) + min
					m["column_to"] = rand.IntN(max-min+1) + min

					if allBoard[m["line_from"]][m["column_from"]] == "#" ||
						allBoard[m["line_from"]][m["column_from"]] == " " {

						movec++
						continue

					}

					//makeMove(chessboard, m)
					h := makeMove(allBoard, m, &newMove, g)
					newMove.History = append(newMove.History, h)

					//printAll(players, allBoard)
					board.SetAllBoard(allBoard)
					g.game.SetGameBoardArr(allBoard)
					g.repo.SaveObject(g.game.GetGameBoard(), 0)
					dataChan <- g

					time.Sleep(2 * time.Second)
				}

				//printHistory(newMove.History)

			} else if newMove.Move != "" {

				fromTo := strings.Split(newMove.Move, "_")

				err := common.CheckSliceLen(fromTo, 2)
				if err != nil {
					fmt.Println("Ошибка ввода:", err.Error())
					return
				}

				from := strings.Split(fromTo[0], ":")
				to := strings.Split(fromTo[1], ":")

				err_from := common.CheckSliceLen(from, 2)
				err_to := common.CheckSliceLen(to, 2)
				if err_from != nil || err_to != nil {
					fmt.Println("Ошибка ввода:", err_from.Error(), err_to.Error())
					return
				}

				from_line_int, err := strconv.Atoi(from[1])
				to_line_int, err := strconv.Atoi(to[1])

				from_column_int := slices.Index(allBoard[0], from[0])
				to_column_int := slices.Index(allBoard[0], to[0])

				if from_column_int < 0 || to_column_int < 0 {
					fmt.Println("Ошибка ввода: Введено некорректное значение координат.")
					return
				}

				m := map[string]int{
					"line_from":   from_line_int,
					"column_from": from_column_int,
					"line_to":     to_line_int,
					"column_to":   to_column_int,
				}

				makeMove(allBoard, m, &newMove, g)
				//g.repo.SaveObject(allBoard, 0)

				board.SetAllBoard(allBoard)
				g.game.SetGameBoardArr(allBoard)
				g.repo.SaveObject(g.game.GetGameBoard(), 0)
				dataChan <- g

				//printAll(player, allBoard)
				//printHistory(newMove.History)
			}
		}

		if newMove.Cancel == true {
			stop <- struct{}{}
			break
		} // Обязательный выход, иначе программа зависнет

	}
}

func setCommands(player string, newMove *models.MoveStruct, playerCount int) /*(bool, int, string)*/ {

	var move string
	var movecount int

	items := []string{"Ввести ход (формат B:4_F:2)", "Автоход (ввести кол-во ходов)", "Сдался"}

	// Создаем меню выбора
	prompt := promptui.Select{
		Label: "Игрок " + player + " выберите действие", // Заголовок меню
		Items: items,
	}

	// Запускаем меню и получаем индекс и текст выбора
	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Ошибка при выборе: %v\n", err)
	}

	// Выводим результат в зависимости от выбора
	switch result {
	case "Автоход (ввести кол-во ходов)":
		fmt.Print("Введите количество ходов (число): ")
		fmt.Scan(&movecount)
		newMove.SetMove(movecount, "", false, playerCount)
	case "Сдался":
		fmt.Println("Игрок " + player + " Сдался")
		newMove.SetMove(0, "", true, playerCount)
	case "Ввести ход (формат B:4_F:2)":
		fmt.Print("Введите ход (формат B:4_F:2): ")
		fmt.Scan(&move)
		newMove.SetMove(0, move, false, playerCount)
	}

}

func printHistory(history []string) {
	for i := range len(history) {
		fmt.Println(history[i])
	}
}

func (g *GameService) PrintAll() {

	common.ClearTerminal()

	g.game.GetGamePlayers().PrintPlayer(0)

	board := g.game.GetGameBoard()
	common.PrintSliceNew(board.GetAllBoard())

	g.game.GetGamePlayers().PrintPlayer(1)
}

func printBoardSlice(player *models.Players, allBoard [][]string) {

	common.ClearTerminal()

	player.PrintPlayer(0)

	common.PrintSliceNew(allBoard)
	player.PrintPlayer(1)
}

func makeMove(allBoard [][]string, m map[string]int,
	move *models.MoveStruct, g *GameService) string {

	g.repo.SaveObject(move, 0)

	//allBoard := chessboard.GetAllBoard()
	allBoard[m["line_to"]][m["column_to"]] =
		allBoard[m["line_from"]][m["column_from"]]

	if (m["line_from"]+m["column_from"])%2 == 0 {
		allBoard[m["line_from"]][m["column_from"]] = "#"

	} else {
		allBoard[m["line_from"]][m["column_from"]] = " "
	}

	h := "Ход " + allBoard[m["0"]][m["column_from"]] + strconv.Itoa(m["line_from"]) + " : " + allBoard[m["0"]][m["column_to"]] + strconv.Itoa(m["line_to"])
	return h
}
