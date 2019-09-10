package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	time.Sleep(3*time.Second)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func t()  {
	s := []int{7,2,8,-9,4,0}
	c := make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x , y := <-c, <-c
	fmt.Println(x,y,x+y)
}
