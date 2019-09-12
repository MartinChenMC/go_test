package main

import (
	"fmt"
	"time"
)

func doWork(ch chan bool) {
	fmt.Println("doWork start")
	time.Sleep(time.Second * 3)
	ch <- true
	fmt.Println("doWork end")
}

func main() {
	ch := make(chan bool)
	go doWork(ch)

	<-ch
}
