package common

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

func CheckValue(a any) error {

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

func CheckSliceLen(s []string, l int) error {

	sl_len := len(s)

	if sl_len != l {

		return errors.New("Введено некорректное значение")

	} else {
		return nil
	}

}

func CheckContains(s []string, symbol string) string {

	var sym string

	if slices.Contains(s, symbol) {
		symbol = symbol + symbol
		sym = CheckContains(s, symbol)

	} else {

		sym = symbol
	}
	return sym

}

func PrintSlice(arr []string, i *int) {

	str := strings.Join(arr, " | ")
	if i == nil {
		fmt.Printf("  %s\n", str)
	} else {
		fmt.Printf("%d %s\n", *i+1, str)
	}

}

func PrintSliceNew(arr [][]string) {

	//fmt.Print("\033[2J")
	stringbuilder := strings.Builder{}

	for j := range len(arr) {
		str := strings.Join(arr[j], " | ") + "\n"
		stringbuilder.Write([]byte(str))

	}

	str_all := stringbuilder.String()
	fmt.Printf("%s", str_all)

}

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
