package main

import "fmt"

func fibonacci1(c, quit chan int) {
	fmt.Println("B")
	x, y := 0, 1
	for {
		fmt.Println("c")
		select {
		case c <- x:
			fmt.Println("d")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("A")
		for i := 0; i < 2; i++ {
			fmt.Println("a")
			fmt.Println(<-c)
			fmt.Println("b")
		}
		quit <- 0
	}()
	fibonacci1(c, quit)
}
