package main

import "fmt"

func main() {

	// Создание буферизированного канала:

	bufferedChannel := make(chan int, 2)

	// Отправка данных в канал:
	bufferedChannel <- 1
	bufferedChannel <- 2

	// Получение данных из канала:
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	// Ещё одна отправка данных:
	bufferedChannel <- 3

	// И ещё одно получение данных:
	fmt.Println(<-bufferedChannel)

	slice := []int{4, 5, 6, 7}

	// Горутина для отправки данных из слайса:
	go func([]int) {
		for _, v := range slice {
			bufferedChannel <- v
		}
		close(bufferedChannel)
	}(slice)

	// Цикл для получения данных:
	for v := range bufferedChannel {
		fmt.Println(v)
	}

}
