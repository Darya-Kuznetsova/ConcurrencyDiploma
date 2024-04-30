package main

import "fmt"

func main() {

	fg := gopherPutBooks()
	sg := gopherDrivesBooks(fg)
	tg := gopherBurnsBooks(sg)
	fog := gopherReturnsCart(tg)

	fmt.Println(<-fog)
}

// Первый гофер, который грузит книги в тележку:

func gopherPutBooks() chan string {
	c1 := make(chan string)
	go func() {
		c1 <- "I'm putting the books"
	}()
	// Канал, который будет использован в следующей функции:
	return c1
}

// Второй гофер, который везёт книги
// В качестве параметра, функция принимает канал, который вернула gopherPutBooks()

func gopherDrivesBooks(firstChannel chan string) chan string {
	c2 := make(chan string)
	fmt.Println(<-firstChannel)
	go func() {
		c2 <- "I'm driving the books"
	}()

	// Канал, который будет использован в следующей функции:
	return c2
}

// Третий гофер, который сжигает книги
// В качестве параметра, функция принимает канал, который вернула gopherDrivesBooks()

func gopherBurnsBooks(secondChannel chan string) chan string {
	c3 := make(chan string)
	fmt.Println(<-secondChannel)
	go func() {
		c3 <- "I'm burning the books"
	}()
	// Канал, который будет использован в следующей функции:
	return c3
}

// Четвертый гофер, который возвращает тележку
// В качестве параметра, функция принимает канал, который вернула gopherBurnsBooks()

func gopherReturnsCart(thirdChannel chan string) chan string {
	c4 := make(chan string)
	fmt.Println(<-thirdChannel)
	go func() {
		c4 <- "I'm returning the cart"
	}()
	return c4
}
