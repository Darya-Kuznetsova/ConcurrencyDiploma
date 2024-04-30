package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4}
	v1 := workOut(slc)
	v2 := fanOut(v1)

	result1 := fanOut(v2)
	result2 := fanOut(v2)

	for range slc {
		select {
		case value := <-result1:
			fmt.Println("1:", value)
		case value := <-result2:
			fmt.Println("2:", value)
		}
	}

}

// Функция принимает слайс с числами:
func workOut(slc []int) <-chan int {

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

// Функция принимает канал, получившийся в результате
// работы функции workOut():
func fanOut(input <-chan int) <-chan int {
	new := make(chan int)

	// Отправляем данные из канала функции workOut()
	// в новый канал:
	go func() {
		defer close(new)

		for data := range input {
			new <- data
		}
	}()
	// Возвращаем новый канал:
	return new
}
