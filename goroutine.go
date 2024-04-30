package main

import (
	"fmt"
	"time"
)

func main() {
	Sum(9, 8)
	go Sum(10, 5)
	time.Sleep(50 * time.Millisecond)
	fmt.Println("END")
}

func Sum(x, y int) {
	fmt.Println(x + y)
}
