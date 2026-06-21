package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Тест мердж реквест 1")

	_ = gin.Default()

	var name string
	fmt.Print("Введите ваше имя: ")
	// Передаем указатель на переменную для записи
	fmt.Scan(&name)
	fmt.Printf("Привет, %s!\n", name)

}
