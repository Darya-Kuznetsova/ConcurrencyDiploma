package main

import (
	"fmt"
	"sync"
)

func main() {
	WaitGroup()
}

func WaitGroup() {
	// Создадим переменную wg с типом sync.WaitGroup:
	var wg sync.WaitGroup

	// Вызовем метод Add(), и передадим 1, так как
	// будем ожидать выполнение одной горутины:
	wg.Add(1)

	go func(x, y int) {
		// Сообщаем о выполнении горутины:
		defer wg.Done()

		fmt.Println(x + y)
	}(5, 10)

	// Ожидает, пока не останется невыполненных горутин:
	wg.Wait()

}
