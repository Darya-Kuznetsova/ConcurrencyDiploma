package main

import (
	"fmt"
	"sync"
)

func main() {

	v1 := work([]int{8, 9, 10, 11})
	v2 := work([]int{12, 13, 14, 15, 16})

	v3 := fanIn(v1, v2)
	for v := range v3 {
		fmt.Println(v)
	}
}

// Функция принимает слайс с числами:
func work(slc []int) <-chan int {

	ch := make(chan int)

	// Горутина, которая отправляет числа в канал:
	go func() {
		defer close(ch)
		for _, v := range slc {
			ch <- v
		}
	}()
	//Возвращаем канал:
	return ch
}

func fanIn(input ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	// Добавляем в WaitGroup число результатов выполнения work():
	wg.Add(len(input))

	// Отправляем данные из канала, полученного в результате выполнения
	// work() в новый канал:
	for _, v := range input {
		go func(ch <-chan int) {
			for {
				value, ok := <-ch

				if !ok {
					wg.Done()
					break
				}

				result <- value
			}
		}(v)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	// Возвращаем итоговый канал:
	return result

}
