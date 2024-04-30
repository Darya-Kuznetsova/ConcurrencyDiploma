package main

import (
	"fmt"
	"sync"
)

func main() {
	GoroutinesCounter()
	GoroutinesCounterRW()
}

func GoroutinesCounter() {
	// Создаём переменную группы ожидания:
	var wg sync.WaitGroup

	// Задаём счётчик:
	var counter int

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
}

func GoroutinesCounterRW() {
	// Создаём переменную группы ожидания:
	var wg sync.WaitGroup

	// Задаём счётчик:
	var counter int

	// Задаём переменную RWMutex:
	var mtx sync.RWMutex

	// Добавляем 1000 горутин в группу ожидания
	wg.Add(2000)
	for i := 0; i < 1000; i++ {
		// Горутина для чтения:
		go func() {
			defer wg.Done()

			// Не блокируем чтение от всех горутин:
			mtx.RLock()
			_ = counter
			mtx.RUnlock()
		}()
		// Горутина для записи в счётчик:
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
}
