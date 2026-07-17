package players

import (
	"fmt"
	common "main/dz6/internal/service/common"
)

/*type Player interface {
	GetPlayer(int) string
	PrintPlayer(int)
	SetPlayers(int)
}*/

type Players struct {
	player1 string
	player2 string
}

func NewPlayers() *Players {
	return &Players{
		player1: "",
		player2: "",
	}
}

// геттер
func (p Players) GetPlayer(i int) string {

	switch i {
	case 0:
		return p.player1
	case 1:
		return p.player2
	default:
		return ""
	}

}

func (p *Players) SetPlayers(count int) {

	var strp string

	for i := range count {

		fmt.Printf("Введите ФИО игрока № %d: ", i+1)
		//fmt.Println("Введите ФИО игрока № ", i+1)
		fmt.Scan(&strp)
		p.SetPlayer(i, strp)

		err := common.CheckValue(p.GetPlayer(i))

		if err != nil {
			fmt.Println("Ошибка ввода:", err.Error())
			return
		}
	}
}

// сеттер
func (p *Players) SetPlayer(i int, s string) {

	switch i {
	case 0:
		p.player1 = s
	case 1:
		p.player2 = s
	}

}

// вывести
func (p Players) PrintPlayer(i int) {

	switch i {
	case 0:
		fmt.Printf("%s\n", p.player1)
	case 1:
		fmt.Printf("%s\n", p.player2)
	}

}
