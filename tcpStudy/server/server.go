package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:6789")
	if err != nil {
		fmt.Println("server start error,", err)
	}
	for true {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("server accept error,", err)
		}
		go readConnection(accept)
	}
}

func readConnection(connection net.Conn) {
	var bytes = make([]byte, 100)
	for true {
		read, err := connection.Read(bytes)
		if err != nil {
			fmt.Println("server read error,", err)
			err := connection.Close()
			if err != nil {
				fmt.Println("关闭连接,一般不会出错吧,反正出错了我也不会处理的")
				return
			}
			return
		}
		fmt.Printf("读取了 [%d] 个字节\t", read)
		fmt.Printf("%s\n", bytes[:read])
	}
}
