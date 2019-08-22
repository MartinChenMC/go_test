package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // 设置2分钟超时
	request := make([]byte, 128)                          // 退出前关闭连接
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		get_str := strings.TrimSpace(string(request[:read_len]))
		fmt.Println(get_str);
		//if read_len == 0 {
		//	break // 客户端已关闭连接
		//} else if get_str == "timestamp" {
		//	daytime := strconv.FormatInt(time.Now().Unix(), 10)
		//	fmt.Println(daytime);
		//	conn.Write([]byte(daytime))
		//} else {
		//	daytime := time.Now().String()
		//	conn.Write([]byte(daytime))
		//}
		conn.Write([]byte("你好 "+get_str))
		request = make([]byte, 128) // 清除上次读取的内容
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
