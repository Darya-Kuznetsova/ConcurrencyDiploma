package main

import "fmt"

func main() {

	buffChannel := make(chan int, 3)
	unbuffChannel := make(chan int)

	select {
	case buffChannel <- 5:
		fmt.Println("First case")
	case unbuffChannel <- 7:
		fmt.Println("Second case")
	}

	select {
	case buffChannel <- 6:
		fmt.Println("First case")
	case buffChannel <- 7:
		fmt.Println("Second case")

	}

	go func() {
		fmt.Println(<-unbuffChannel)
	}()

	select {
	case unbuffChannel <- 8:
		fmt.Println("First case")
	default:
		fmt.Println("Default")

	}
	select {
	case <-buffChannel:
		fmt.Println("buff")
	case <-unbuffChannel:
		fmt.Println("unbuff")
	}

}
