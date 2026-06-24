package main

import (
	"errors"
	"fmt"
	printlines "main/dz3/packages"
	"strings"
)

func main() {

	// инициализация
	var length int
	var str string
	var player1 string
	var player2 string
	countflag := true

	//  ввод значений переменных
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&length)

	err := checkValue(length)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	fmt.Print("Введите ФИО игрока №1: ")
	fmt.Scan(&player1)

	err = checkValue(player1)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	fmt.Print("Введите ФИО игрока №2: ")
	fmt.Scan(&player2)

	err = checkValue(player2)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
		return
	}

	// формирование и вывод линий
	fmt.Printf("%s\n", player1)

	letters := printlines.PrintLetters(length)
	fmt.Printf("  %s\n", letters)

	for i := 0; i < length; i++ {

		if i <= 1 {
			//  расставляем белые
			str = printlines.PrintWhite(length)
		} else if length > 4 && i > 1 && i < length-2 {
			//  рисуем поле
			countflag, str = cycleJ(countflag, str)
			str = printlines.Printline(length, str)
		} else {
			//  расставляем черные
			str = printlines.PrintBlack(length)
		}

		fmt.Printf("%d %s\n", i+1, str)
	}

	fmt.Printf("%s\n", player2)

}

func cycleJ(countflag bool, str string) (bool, string) {

	if countflag {
		str = " "
		countflag = false
	} else {

		str = "#"
		countflag = true
	}
	return countflag, str
}

func checkValue(a any) error {

	b, ok_int := a.(int)
	if ok_int == true {

		if b == 0 {

			return errors.New("Длина доски должна быть больше нуля")

		} else {
			return nil
		}

	}

	str, ok_string := a.(string)
	if ok_string == true {

		if strings.TrimSpace(str) == "" {

			return errors.New("Значение ФИО игрока не введено")

		} else {
			return nil
		}

	}

	return errors.New("Тип не определен")
}
