package main

import "fmt"

func main() {

	// Создаём небуферизированный канал с типом int:

	unbufferedChannel := make(chan int)

	// Горутина для отправки данных в канал:
	go func() {
		unbufferedChannel <- 8
	}()

	// Получение данных из канала:
	fmt.Println(<-unbufferedChannel)

	// Горутина для получения данных из канала:
	go func() {
		fmt.Println(<-unbufferedChannel)
	}()

	// Отправка данных в канал:
	unbufferedChannel <- 8

	// Создаём небуферизированный канал типа bool:

	boolChannel := make(chan bool)

	// Запускаем Sum как горутины:

	go Sum2(9, 8, boolChannel)
	go Sum2(10, 5, boolChannel)

	// Отправляем данные из канала в main()
	<-boolChannel

}

func Sum2(x, y int, c chan bool) {
	fmt.Println(x + y)
	c <- true
}
