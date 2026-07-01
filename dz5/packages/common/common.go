package common

import (
	"errors"
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
