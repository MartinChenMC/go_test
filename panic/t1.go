package main

import "fmt"

func main() {
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")

	panic("555")
}
