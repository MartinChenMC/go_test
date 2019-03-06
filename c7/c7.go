// +build go1.10

package main

//void SayHello(_GoString_  s);
import "C"

import "fmt"

func main()  {
	C.SayHello("hello,World Martin \n")

	pid := unix.Getpid()
	fmt.Println("main Getpid", pid)
}

//export SayHell
func SayHello(s string)  {
	fmt.Print(s)

	pid := unix.Getpid()
	fmt.Println("SayHello Getpid", pid)
}