package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	GoroutinesCounterTime()
	GoroutinesCounterAtomic()
}

func GoroutinesCounterTime() {
	// Добавляем время старта программы:
	start := time.Now()
	// Создаём переменную группы ожидания:
	var wg sync.WaitGroup

	// Задаём счётчик:
	var counter int64

	// Задаём переменную Mutex:
	var mtx sync.Mutex

	// Добавляем 1000 горутин в группу ожидания
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Сообщаем о выполнении горутины:
			defer wg.Done()

			// Увеличиваем счётчик
			// и блокируем доступ к counter для других горутин:
			mtx.Lock()
			counter++

			// Разблокируем доступ:
			mtx.Unlock()
		}()
	}
	// Ожидаем, пока все горутины будут выполнены:
	wg.Wait()

	// Выводим на экран счётчик:
	fmt.Println(counter)
	// Выводим на экран время, за которое была выполнена программа:
	fmt.Println("Time:", time.Now().Sub(start).Seconds())
}

func GoroutinesCounterAtomic() {
	// Добавляем время старта программы:
	start := time.Now()

	// Создаём переменную группы ожидания:
	var wg sync.WaitGroup

	// Задаём счётчик:
	var counter int64

	// Добавляем 1000 горутин в группу ожидания
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Сообщаем о выполнении горутины:
			defer wg.Done()
			// Добавляем значение 1 в counter:
			atomic.AddInt64(&counter, 1)

		}()
	}
	// Ожидаем, пока все горутины будут выполнены:
	wg.Wait()

	// Выводим на экран счётчик:
	fmt.Println(counter)

	// Выводим на экран время, за которое была выполнена программа:
	fmt.Println("Time:", time.Now().Sub(start).Seconds())
}
