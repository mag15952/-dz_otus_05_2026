package service

import (
	"fmt"
	"log"
	"main/dz7/internal/models"
	"main/dz7/internal/repository"
	"main/dz7/internal/service"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func setCommandsNew(player string, newMove *models.MoveStruct, boardcount int) /*(bool, int, string)*/ {

	var move string
	var movecount int

	items := []string{"Ввести ход (формат B:4_F:2)", "Автоход (ввести кол-во ходов)", "Сдался"}
	if boardcount > 1 {
		items = []string{"Автоход (ввести кол-во ходов)", "Сдался"}
	}

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
		newMove.SetMove(movecount, "", false)
	case "Сдался":
		fmt.Println("Игрок " + player + " Сдался")
		newMove.SetMove(0, "", true)
	case "Ввести ход (формат B:4_F:2)":
		fmt.Print("Введите ход (формат B:4_F:2): ")
		fmt.Scan(&move)
		newMove.SetMove(0, move, false)
	}

}

func PlayChessNew(player *models.Players, length int, chessboard *models.Chessboard) {

	for {

		var newMove models.MoveStruct
		newMove = models.NewMove()

		for i := range 2 {

			p := player.GetPlayer(i)
			//cancel, movecount, move = setCommands(p)
			setCommandsNew(p, &newMove, chessboard.Boardcount)

			if newMove.Cancel == true {
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

					if chessboard.AllBoard[m["line_from"]][m["column_from"]] == "#" ||
						chessboard.AllBoard[m["line_from"]][m["column_from"]] == " " {

						movec++
						continue

					}

					//makeMove(chessboard, m)
					h := makeMove(chessboard, m)
					newMove.History = append(newMove.History, h)

					printAllAgain(player, chessboard)

					time.Sleep(2 * time.Second)
				}

				printHistory(newMove.History)

			} else if newMove.Move != "" {

				fromTo := strings.Split(newMove.Move, "_")

				err := service.CheckSliceLen(fromTo, 2)
				if err != nil {
					fmt.Println("Ошибка ввода:", err.Error())
					return
				}

				from := strings.Split(fromTo[0], ":")
				to := strings.Split(fromTo[1], ":")

				err_from := service.CheckSliceLen(from, 2)
				err_to := service.CheckSliceLen(to, 2)
				if err_from != nil || err_to != nil {
					fmt.Println("Ошибка ввода:", err_from.Error(), err_to.Error())
					return
				}

				from_line_int, err := strconv.Atoi(from[1])
				to_line_int, err := strconv.Atoi(to[1])

				from_column_int := slices.Index(chessboard.AllBoard[0], from[0])
				to_column_int := slices.Index(chessboard.AllBoard[0], to[0])

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

				makeMove(chessboard, m)
				printAllAgain(player, chessboard)
				printHistory(newMove.History)
			}
		}

		if newMove.Cancel == true {
			break
		} // Обязательный выход, иначе программа зависнет
	}
}

func printHistory(history []string) {
	for i := range len(history) {
		fmt.Println(history[i])
	}
}

func printAllAgain(player *models.Players, chessboard *models.Chessboard) {

	service.ClearTerminal()

	repository.PrintObject(player, 0)

	service.PrintSliceNew(chessboard.AllBoard)

	repository.PrintObject(player, 1)
}

func makeMove(chessboard *models.Chessboard, m map[string]int) string {

	chessboard.AllBoard[m["line_to"]][m["column_to"]] =
		chessboard.AllBoard[m["line_from"]][m["column_from"]]

	if (m["line_from"]+m["column_from"])%2 == 0 {
		chessboard.AllBoard[m["line_from"]][m["column_from"]] = "#"

	} else {
		chessboard.AllBoard[m["line_from"]][m["column_from"]] = " "
	}

	h := "Ход " + chessboard.AllBoard[m["0"]][m["column_from"]] + strconv.Itoa(m["line_from"]) + " : " + chessboard.AllBoard[m["0"]][m["column_to"]] + strconv.Itoa(m["line_to"])
	return h
}
