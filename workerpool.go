package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	WorkerPool()
}

func WorkerPool() {
	// Создаём контекст с таймаутом:
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	wg := &sync.WaitGroup{}

	// Создаём канал numbers для получения данных и results для отправки:
	numbers, results := make(chan int, 5), make(chan int, 5)

	// Запускаем воркеры:
	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbers, results)
		}()
	}
	// Наполняем канал numbers:
	go func() {
		for i := 0; i < 100; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	// Ожидаем выполнения работы воркеров и закрываем result канал:
	go func() {
		wg.Wait()
		close(results)
	}()
	// Счётчик результатов, которые успели выполниться:
	var counter int
	for resultValue := range results {
		counter++
		fmt.Println(resultValue)
	}
	fmt.Println(counter)
}

// Создание воркера:

func worker(ctx context.Context, Numbers <-chan int, Result chan<- int) {
	for {
		select {
		// Вариант, когда контекст отменён:
		case <-ctx.Done():
			return
		// Проверяем наличие значений в канале и
		// канал отправляет их в переменную:
		case value, ok := <-Numbers:
			if !ok {
				return
			}
			// Имитируем затрату времени и
			// канал получает квадрат числа
			time.Sleep(time.Millisecond)
			Result <- value * value
		}
	}
}
