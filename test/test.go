package main

import (
	"fmt"
	"time"
)

type Сh struct {
	length   int
	allBoard [][]string
	Chbd     chan [][]string
}

func main() {

	dataChan := make(chan string)

	// Создаем тикер на 1 секунду
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Обязательно останавливаем при выходе

	var currentData = "Ожидание данных..."

	// Горутина-эмулятор, которая шлет данные раз в 3 секунды
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(3 * time.Second)
			dataChan <- fmt.Sprintf("Данные №%d", i)
		}
	}()

	// Главный цикл отрисовки и обработки событий
	for {
		select {
		case data := <-dataChan:
			// Обновляем текущее состояние при получении новых данных
			currentData = data

		case <-ticker.C:
			// Каждую секунду перерисовываем экран
			printTerminal(currentData)
		}
	}
}

func printTerminal(data string) {
	// Очистка экрана и возврат курсора в начало (ANSI-последовательность)
	fmt.Print("\033[H\033[2J")
	fmt.Println("=== ТЕРМИНАЛ (обновлено раз в сек) ===")
	fmt.Println("Статус:", data)
	fmt.Println("Время:", time.Now().Format("15:04:05"))
}
