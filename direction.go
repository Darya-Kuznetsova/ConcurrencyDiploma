package main

import "fmt"

func main() {

	// Создаём небуферизированный канал:
	unbuffChannel := make(chan int)

	// Горутина, которая принимает канал только для получения данных:
	go func(receiveChan chan<- int) {
		receiveChan <- 8
	}(unbuffChannel)

	// Отправляем данные из канала:
	fmt.Println(<-unbuffChannel)

	// Грутина, которая принимает канал только для отправки данных:
	go func(sendChannel <-chan int) {
		fmt.Println(<-sendChannel)
	}(unbuffChannel)

	// Получение данных:
	unbuffChannel <- 9
}
