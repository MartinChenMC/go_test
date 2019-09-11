package main

import "fmt"

func main()  {
	ch := make(chan int,2)
	fmt.Println(len(ch))
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
}
