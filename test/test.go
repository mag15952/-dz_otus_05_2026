package main

import (
	"fmt"
	"time"
)

func main() {

	// Версия с очисткой и абсолютным позиционированием
	// fmt.Print("\033[H\033[J")
	for i := 0; i < 3; i++ {
		fmt.Println("---")
	}

	toggle := true
	for {
		// Версия с очисткой и абсолютным позиционированием
		// fmt.Printf("\033[%d;%dH", 2, 2)

		// Версия без очистки
		fmt.Printf("\033[%dA", 2)
		fmt.Printf("\033[%dC", 1)

		if toggle {
			fmt.Print("x")
		} else {
			fmt.Print("-")
		}

		// Версия с очисткой и абсолютным позиционированием
		// fmt.Printf("\033[%d;%dH", 4, 1)

		// Версия без очистки
		fmt.Printf("\033[%dB", 2)
		fmt.Printf("\033[%dD", 2)

		toggle = !toggle
		time.Sleep(time.Second)
	}

	/*for i := 0; i <= 10; i++ {
		// \r возвращает каретку, чтобы перезаписать текущую строку

		fmt.Printf("\rПрогресс: %d%%", i)

		time.Sleep(2 * time.Second)

		fmt.Printf("\rПрогресс222: %d%%", i)

		time.Sleep(2 * time.Second)
	}
	fmt.Println("\nГотово!")*/

	//fmt.Println("yFDLFKJLKSJ\n")

	//fmt.Print("\033[2J")

	/*for i := range 2 {
		// Возвращаем курсор в начало экрана перед новым выводом
		fmt.Printf("\033[HСчетчик: %d/10", i)
		time.Sleep(1 * time.Second)
		fmt.Printf("\033[HСчетчик: %d/5", i)
		time.Sleep(1 * time.Second)
	}*/
	//fmt.Println(rand.IntN(2))

	//fmt.Print("\033[2J")

	//fmt.Println("yFDLFKJLKSJ\n")

}
