package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	//_, err = conn.Write([]byte("timestamp"))
	//checkError(err)
	//fmt.Println("here")
	//time.Sleep(time.Second*2)
	//result, err := readFully(conn)
	//checkError(err)
	//fmt.Println(string(result))
	//os.Exit(0)
	for {
		sendMsg(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func sendMsg(conn net.Conn)  {
	var input string
	_, err := fmt.Scanf("%s\r\n", &input)
	fmt.Println(input);
	checkError(err)
	if (input == "26") {
		os.Exit(0)
	}
	_, err = conn.Write([]byte(input))
	checkError(err)

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	result.Write(buf[0:n])
	fmt.Println(string(result.Bytes()))
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		fmt.Println(string(result.Bytes()))
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}