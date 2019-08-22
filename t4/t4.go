package main

import (
	"fmt"
	"time"
)

func main() {
	go getNum()
	go getNum()


	time.Sleep(time.Duration(3)*time.Second)
}

func getNum() {
	timeUnix:=time.Now().Unix()
	fmt.Println(timeUnix)
	time.Sleep(time.Second)
}